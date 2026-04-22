package pubsub

type Broker struct {
	ch          chan Event
	subscribers map[chan Event]struct{}
}

func NewBrocker() *Broker {
	ch := make(chan Event)
	subscribers := make(map[chan Event]struct{})
	return &Broker{
		ch:          ch,
		subscribers: subscribers,
	}
}

func (b *Broker) Subscribe(ch chan Event) {
	b.subscribers[ch] = struct{}{}
}

func (b *Broker) Publish(e Event) {

	// チャンネルに通知を送る
	for ch := range b.subscribers {
		ch <- e
	}
}

// 全てのサブスクライバーのチャンネルをcloseする
func (b *Broker) Close() {
	for ch := range b.subscribers {
		close(ch)
	}
}
