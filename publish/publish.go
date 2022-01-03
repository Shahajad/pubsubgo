package publish

import (
	"pubsubgo/pubsub"
)

func CreateTopic(TopicID string) {
	pubsub.CreateTopic(TopicID)
}
func DeleteTopic(TopicID string) {
	pubsub.DeleteTopic(TopicID)
}

func Publish(TopicId string, message string) {
	go pubsub.Publish(TopicId, message)
}
