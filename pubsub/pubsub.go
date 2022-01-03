package pubsub

import (
	"fmt"
	"math/rand"
	"sync"
)

type DataEvent struct {
	Data    interface{}
	Message string
}

type DataChannel chan DataEvent

type EventBus struct {
	topics                  []string
	subscriptionIds         []string
	topicSubscriptionMap    map[string][]string
	subscriptionIdChanelMap map[string]DataChannel
	subscriptionIdFuncMap   map[string]func(data DataEvent)
	rm                      sync.RWMutex
}

var eb = &EventBus{
	topics:                  []string{},
	subscriptionIds:         []string{},
	topicSubscriptionMap:    map[string][]string{},
	subscriptionIdChanelMap: map[string]DataChannel{},
	subscriptionIdFuncMap:   map[string]func(data DataEvent){},
}

func CreateTopic(TopicID string) bool {
	index := GetIndex(eb.topics, TopicID)
	if index >= 0 {
		// don't create topic if topic already exist
		return false
	}
	eb.topics = append(eb.topics, TopicID)
	eb.topicSubscriptionMap[TopicID] = []string{}
	return true
}
func DeleteTopic(TopicID string) bool {
	index := GetIndex(eb.topics, TopicID)
	if index < 0 {
		return false
	}
	eb.topics = RemoveIndex(eb.topics, index)
	for _, subscriberId := range eb.subscriptionIds {
		DeleteSubscription(subscriberId)
	}
	delete(eb.topicSubscriptionMap, TopicID)
	return true
}

func AddSubscription(TopicID string, SubscriptionID string) bool {
	index := GetIndex(eb.topics, TopicID)
	if index < 0 {
		// don't add subscription if topic doesn't exist
		return false
	}

	index = GetIndex(eb.subscriptionIds, SubscriptionID)
	if index >= 0 {
		// don't add subscription if SubscriptionID already exist
		return false
	}
	eb.subscriptionIds = append(eb.subscriptionIds, SubscriptionID)
	eb.topicSubscriptionMap[TopicID] = append(eb.topicSubscriptionMap[TopicID], SubscriptionID)
	eb.subscriptionIdChanelMap[SubscriptionID] = make(DataChannel)
	return true
}

func DeleteSubscription(SubscriptionID string) bool {
	eb.rm.Lock()
	index := GetIndex(eb.subscriptionIds, SubscriptionID)
	if index < 0 {
		eb.rm.Unlock()
		return false
	}
	eb.subscriptionIds = RemoveIndex(eb.subscriptionIds, index)
	delete(eb.subscriptionIdChanelMap, SubscriptionID)
	delete(eb.subscriptionIdFuncMap, SubscriptionID)
	fmt.Println("Deleted ", SubscriptionID)
	eb.rm.Unlock()
	return true
}

func Publish(TopicId string, data interface{}) {
	Run()
	eb.rm.Lock()
	if subscriptions, found := eb.topicSubscriptionMap[TopicId]; found {
		subscriptions := append([]string{}, subscriptions...)
		for _, subscriptionId := range subscriptions {
			if _, found := eb.subscriptionIdFuncMap[subscriptionId]; found {
				eb.subscriptionIdChanelMap[subscriptionId] <- DataEvent{Data: data, Message: TopicId}
			}

		}
	}
	eb.rm.Unlock()
}

func Run() {
	for _, subscriberId := range eb.subscriptionIds {
		if chans, found := eb.subscriptionIdChanelMap[subscriberId]; found {
			if fn, found := eb.subscriptionIdFuncMap[subscriberId]; found {
				go sendToSubscriber(chans, subscriberId, fn)
			}
		}
	}
}

func sendToSubscriber(ch chan DataEvent, subscriptionId string, fn func(data DataEvent)) {
	notProcessedCh := make(DataChannel, 10)
	for {
		//if _, found := eb.subscriptionIdFuncMap[subscriptionId]; !found {
		//	break
		//}
		if len(notProcessedCh) > 0 {
			data := <-notProcessedCh
			ack := ConsumeMessage(data, subscriptionId, fn, true)
			if !ack {
				notProcessedCh <- data
			}
		} else {
			select {
			case data := <-ch:
				ack := false
				x := rand.Intn(1000)
				if x%6 == 0 {
					// failing some message (around 1/3rd) to reprocess it again
					// it will give ack as false
					ack = ConsumeMessage(DataEvent{Data: "", Message: ""}, subscriptionId, fn, false)
				} else {
					ack = ConsumeMessage(data, subscriptionId, fn, false)
				}

				if !ack {
					notProcessedCh <- data
					break
				}
			default:
				//time.Sleep(time.Second * 1)
				break
			}
		}
	}
	//fmt.Println(" break", subscriptionId)
}

func UnSubscribe(SubscriptionID string) {
	eb.rm.Lock()
	delete(eb.subscriptionIdFuncMap, SubscriptionID)
	fmt.Println("Unsubscribed ", SubscriptionID)
	eb.rm.Unlock()

}

func Subscribe(SubscriptionID string, SubscriberFunc func(data DataEvent)) bool {
	if _, found := eb.subscriptionIdFuncMap[SubscriptionID]; !found {
		eb.subscriptionIdFuncMap[SubscriptionID] = SubscriberFunc
		return true
	}
	return false
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func GetIndex(s []string, item string) int {
	index := -1
	for i := 0; i < len(s); i++ {
		if item == s[i] {
			index = i
			return index
		}
	}
	return index
}

func Ack(SubscriptionID string, MessageID string) bool {
	if SubscriptionID != "" && MessageID != "" {
		return true
	}
	return false
}

func ConsumeMessage(data DataEvent, subscriptionId string, callBackFn func(data DataEvent), retry bool) bool {
	if retry {
		fmt.Println("retry consuming message for subscriber: ", subscriptionId, " and message: ", data)
		callBackFn(data)
		fmt.Println()
		return Ack(subscriptionId, data.Message)
	} else {
		fmt.Println("consuming message for subscriber: ", subscriptionId, " and message: ", data)
		callBackFn(data)
		fmt.Println()
		return Ack(subscriptionId, data.Message)
	}

}
