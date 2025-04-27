package pubsub

type Message struct {
	Topic string
	Body  string
}

func NewMessage(msg, topic string) *Message {
	return &Message{Topic: topic, Body: msg}
}
