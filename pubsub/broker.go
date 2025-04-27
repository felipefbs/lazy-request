package pubsub

import "sync"

type Subscribers map[string]*Subscriber

type Broker struct {
	subscribers Subscribers
	topics      map[string]Subscribers
	mu          sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: Subscribers{},
		topics:      map[string]Subscribers{},
	}
}

func (b *Broker) Subscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.topics[topic] == nil {
		b.topics[topic] = Subscribers{}
	}

	s.AddTopic(topic)
	b.topics[topic][s.id] = s
}

func (b *Broker) Publish(topic, msg string) {
	b.mu.RLock()
	bTopics := b.topics[topic]
	b.mu.RUnlock()

	for _, s := range bTopics {
		m := NewMessage(msg, topic)
		if !s.active {
			return
		}
		go (func(s *Subscriber) {
			s.Signal(m)
		})(s)
	}
}

func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	delete(b.topics[topic], s.id)

	s.RemoveTopic(topic)
}

func (b *Broker) RemoveSubscriber(s *Subscriber) {
	for topic := range s.topics {
		b.Unsubscribe(s, topic)
	}

	b.mu.Lock()
	delete(b.subscribers, s.id)
	b.mu.Unlock()

	s.Destruct()
}

func (b *Broker) Broadcast(msg string, topics []string) {
	for _, topic := range topics {
		for _, s := range b.topics[topic] {
			m := NewMessage(msg, topic)
			go (func(s *Subscriber) {
				s.Signal(m)
			})(s)
		}
	}
}

func (b *Broker) AddSubscriber() *Subscriber {
	// Add subscriber to the broker.
	b.mu.Lock()
	defer b.mu.Unlock()
	id, s := CreateNewSubscriber()
	b.subscribers[id] = s
	return s
}
