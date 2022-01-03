package main

import (
	"fmt"
	"pubsubgo/publish"
	"pubsubgo/pubsub"
	"pubsubgo/subscriber"
	"time"
)

func main() {
	publish.CreateTopic("topic1")
	publish.CreateTopic("topic2")

	subscriber.AddSubscription("topic1", "sub1")
	subscriber.AddSubscription("topic1", "sub2")
	subscriber.AddSubscription("topic2", "sub3")

	subscriber.Subscribe("sub1", hello1)
	subscriber.Subscribe("sub2", hello2)
	subscriber.Subscribe("sub3", hello3)

	publish.Publish("topic1", "hi 1 from")
	publish.Publish("topic1", "hi 2 from")

	time.Sleep(time.Second * 2)

	subscriber.UnSubscribe("sub1")
	subscriber.DeleteSubscription("sub2")
	subscriber.DeleteSubscription("sub2")

	publish.Publish("topic1", "hi 3 from")
	publish.Publish("topic1", "hi 4 from")
	publish.Publish("topic2", "hi 1 from")
	publish.Publish("topic2", "hi 1 from")
	time.Sleep(time.Second * 10)
	fmt.Println("Main Method completed")
}

func hello1(event pubsub.DataEvent) {
	if event.Message == "" {
		fmt.Println("message failed to consumed using function hello1")
	} else {
		fmt.Println("message consumed using function hello1 for ", event)
	}

}
func hello2(event pubsub.DataEvent) {
	if event.Message == "" {
		fmt.Println("message failed to consumed using function hello2")
	} else {
		fmt.Println("message consumed using function hello2 for", event)
	}
}
func hello3(event pubsub.DataEvent) {
	if event.Message == "" {
		fmt.Println("message failed to consumed using function hello3")
	} else {
		fmt.Println("message consumed using function hello3 for", event)
	}
}

func Test() {
	publish.CreateTopic("topic1")
	subscriber.AddSubscription("topic1", "sub1")
	subscriber.Subscribe("sub1", hello1)
	publish.Publish("topic1", "hi 1 from")
	publish.Publish("topic1", "hi 2 from")
	time.Sleep(time.Second * 2)
}

func Test1() {
	publish.CreateTopic("topic1")
	subscriber.AddSubscription("topic1", "sub1")
	subscriber.Subscribe("sub1", hello1)
	subscriber.AddSubscription("topic1", "sub2")
	subscriber.Subscribe("sub2", hello2)
	publish.Publish("topic1", "hi 1 from")
	publish.Publish("topic1", "hi 2 from")
	time.Sleep(time.Second * 3)
}

func Test2() {
	publish.CreateTopic("topic1")
	subscriber.AddSubscription("topic1", "sub1")
	subscriber.Subscribe("sub1", hello1)
	subscriber.AddSubscription("topic1", "sub2")
	subscriber.Subscribe("sub2", hello2)

	publish.Publish("topic1", "hi 1 from")
	publish.Publish("topic1", "hi 2 from")
	time.Sleep(time.Second * 3)
	subscriber.UnSubscribe("sub1")
	publish.Publish("topic1", "hi 3 from")
	publish.Publish("topic1", "hi 4 from")
	time.Sleep(time.Second * 3)
}

func Test3() {
	publish.CreateTopic("topic1")
	publish.CreateTopic("topic2")

	subscriber.AddSubscription("topic1", "sub1")
	subscriber.Subscribe("sub1", hello1)
	subscriber.AddSubscription("topic2", "sub2")
	subscriber.Subscribe("sub2", hello2)

	publish.Publish("topic1", "hi 1 from")
	publish.Publish("topic2", "hi 1 from")
	publish.Publish("topic1", "hi 2 from")
	publish.Publish("topic2", "hi 2 from")
	time.Sleep(time.Second * 3)
}

func Test4() {
	publish.CreateTopic("topic1")
	publish.CreateTopic("topic2")

	subscriber.AddSubscription("topic1", "sub1")
	subscriber.Subscribe("sub1", hello1)

	publish.Publish("topic1", "hi 1 from")

	time.Sleep(time.Second * 1)
	publish.DeleteTopic("topic1")

	publish.Publish("topic1", "hi 2 from")
	time.Sleep(time.Second * 3)
}
