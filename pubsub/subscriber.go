package pubsub

import (
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

type Subscriber struct {
	id       string
	messages chan *Message
	topics   map[string]bool
	active   bool
	mu       sync.RWMutex
}

func CreateNewSubscriber() (string, *Subscriber) {
	b := make([]byte, 8)

	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	id := fmt.Sprintf("%X-%X", b[0:4], b[4:8])

	return id, &Subscriber{
		id:       id,
		messages: make(chan *Message),
		topics:   map[string]bool{},
		active:   true,
	}
}

func (s *Subscriber) Signal(msg *Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.active {
		s.messages <- msg
	}
}

func (s *Subscriber) Destruct() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.active = false
	close(s.messages)
}

func (s *Subscriber) AddTopic(topic string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.topics[topic] = true
}

func (s *Subscriber) RemoveTopic(topic string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	delete(s.topics, topic)
}

func (s *Subscriber) Listen() {
	for {
		if msg, ok := <-s.messages; ok {
			fmt.Printf("Subscriber %s, received: %s from topic: %s\n", s.id, msg.Body, msg.Topic)
		}
	}
}
