package subscriber

import "go-notification/pubsub"

type Subscriber struct {
	br *pubsub.Broker
	s  *pubsub.Subscriber
}

func NewSubscriber(br *pubsub.Broker) *Subscriber {
	return &Subscriber{
		br: br,
	}
}

func (s *Subscriber) Subscribe() *Subscriber {
	s.s = s.br.Subscribe()
	return s
}
