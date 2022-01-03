package subscriber

import (
	"fmt"
	"pubsubgo/pubsub"
	"testing"
)

func TestAddSubscription(t *testing.T) {
	AddSubscription("topic1", "sub1")
}

func TestDeleteSubscription(t *testing.T) {
	DeleteSubscription("sub1")
}

func TestSubscribe(t *testing.T) {

	Subscribe("sub1", test)
}
func TestUnSubscribe(t *testing.T) {
	UnSubscribe("sub1")
}

func test(event pubsub.DataEvent) {
	if event.Message == "" {
		fmt.Println("message failed to consumed using function hello1")
	} else {
		fmt.Println("message consumed using function hello1 for ", event)
	}
}
