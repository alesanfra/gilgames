package broker

import (
	"errors"
	"log"


)

type Broker struct {
	topics map[string][]string
}

func New() Broker {
	return Broker{topics: map[string][]string{}}
}

func (b *Broker) TopicExists(topic string) bool {
	_, exists := b.topics[topic]
	return exists
}

func (b *Broker) CreateTopic (topic string) error {
	b.topics[topic] = []string{}
	log.Println("Created topic:", topic)
	return nil
}

func (b *Broker) SendMessage (topic string, message string) error {
	q, exists := b.topics[topic]
	if !exists {
		return errors.New("Topic not found")
	}

	b.topics[topic]=append(q, message)
	return nil
}