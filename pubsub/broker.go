package pubsub

type Broker struct {
	ch          chan Event
	subscribers map[chan Event]bool
}

func NewBrocker() *Broker {
	ch := make(chan Event)
	subscribers := make(map[chan Event]bool)
	return &Broker{
		ch:          ch,
		subscribers: subscribers,
	}
}

func (b *Broker) Subscribe(ch chan Event) {
	b.subscribers[ch] = true
}

func (b *Broker) Publish(e Event) {

	// チャンネルに通知を送る
	for ch, s := range b.subscribers {
		if s {
			ch <- e
		}

	}
}
