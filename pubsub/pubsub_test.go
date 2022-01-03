package pubsub

import (
	"fmt"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	CreateTopic("topic1")
}

func TestDeleteTopic(t *testing.T) {
	CreateTopic("topic1")
	DeleteTopic("topic1")
}

func TestUnSubscribe(t *testing.T) {
	UnSubscribe("sub1")
}

func TestAddSubscription(t *testing.T) {
	AddSubscription("topic1", "sub1")
}

func TestDeleteSubscription(t *testing.T) {
	AddSubscription("topic1", "sub1")
	DeleteSubscription("sub1")
}

func TestSubscribe(t *testing.T) {
	Subscribe("sub1", hello1)
}

func TestConsumeMessage(t *testing.T) {
	ConsumeMessage(DataEvent{Data: "test", Message: "test"}, "sub1", hello1, false)
}

func TestPublish(t *testing.T) {
	CreateTopic("topic1")
	AddSubscription("topic1", "sub1")
	Subscribe("sub1", hello1)

	Publish("topic1", "test")
}

func TestAck(t *testing.T) {
	Ack("sub1", "test")
}

func hello1(event DataEvent) {
	if event.Message == "" {
		fmt.Println("message failed to consumed using function hello1")
	} else {
		fmt.Println("message consumed using function hello1 for ", event)
	}

}
