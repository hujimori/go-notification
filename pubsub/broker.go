package pubsub

type Broker struct {
	ch          chan Event
	subscribers map[*Subscriber]struct{}
}

func NewBrocker() *Broker {
	ch := make(chan Event)
	subscribers := make(map[*Subscriber]struct{})
	return &Broker{
		ch:          ch,
		subscribers: subscribers,
	}
}

func (b *Broker) Start() {
	go b.run()
}

func (b *Broker) run() {
	for e := range b.ch {
		b.dispatch(e)
	}
	b.closeSubscribers()
}

func (b *Broker) dispatch(e Event) {
	// dispatch：routing + sendingのセット動詞
	// ソフトウェアの文脈では「適切な宛先に振り分けて送る」
	for s := range b.subscribers {
		s.ch <- e
	}
}

func (b *Broker) closeSubscribers() {
	// 全てのチャンネルをcloseする
	for s := range b.subscribers {
		close(s.ch)
	}
}

func (b *Broker) Subscribe() *Subscriber {
	s := &Subscriber{ch: make(chan Event)}
	b.subscribers[s] = struct{}{}
	return s
}

func (b *Broker) Publish(e Event) {
	// チャンネルに通知を送る
	b.ch <- e
}

// 全てのチャンネルをcloseする
func (b *Broker) Close() {
	close(b.ch)
}
