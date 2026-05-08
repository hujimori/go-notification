package publisher

import "go-notification/pubsub"

type Publisher struct {
	br *pubsub.Broker
}

func NewPublisher(br *pubsub.Broker) *Publisher {
	return &Publisher{
		br: br,
	}
}

func (p *Publisher) Publish(e pubsub.Event) {
	p.br.Publish(e)
}
