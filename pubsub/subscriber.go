package pubsub

type Subscriber struct {
	ch chan Event
}

func (s *Subscriber) Events() <-chan Event {
	return s.ch
}
