package subscriber

import (
	"pubsubgo/pubsub"
)

func AddSubscription(TopicID string, SubscriptionID string) bool {
	return pubsub.AddSubscription(TopicID, SubscriptionID)
}

func DeleteSubscription(SubscriptionID string) {
	pubsub.DeleteSubscription(SubscriptionID)
}

func UnSubscribe(SubscriptionID string) {
	pubsub.UnSubscribe(SubscriptionID)
}

func Subscribe(SubscriptionID string, SubscriberFunc func(data pubsub.DataEvent)) {
	pubsub.Subscribe(SubscriptionID, SubscriberFunc)
}
