package main

import "testing"

func TestTest(t *testing.T) {
	// one topic, one subscription created, two message published on same topic
	Test()
}

func TestTest1(t *testing.T) {
	// one topic, two subscription created, two message published on same topic
	Test1()
}
func TestTest2(t *testing.T) {
	// one topic, two subscription/subscriber created, 4 message published on same topic
	// 2 messages were published after unsubscribing "sub1" subscription
	Test2()
}

func TestTest3(t *testing.T) {
	// two topic, two subscription/subscriber created, 2 message each published on both topic
	Test3()
}

func TestTest4(t *testing.T) {
	// one topic, one subscription/subscriber created
	// topic deleted after publishing one message and then published new message
	Test4()
}
