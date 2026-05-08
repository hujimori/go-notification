package pubsub

import (
	"fmt"
	"sync"
)

type Broker struct {
	mu          sync.Mutex
	ch          chan Event
	subscribers map[*Subscriber]struct{}
	done        chan struct{}
}

func NewBrocker() *Broker {
	ch := make(chan Event)
	subscribers := make(map[*Subscriber]struct{})
	done := make(chan struct{})
	return &Broker{
		ch:          ch,
		subscribers: subscribers,
		done:        done,
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
	close(b.done)
}

func (b *Broker) dispatch(e Event) {
	// dispatch：routing + sendingのセット動詞
	// ソフトウェアの文脈では「適切な宛先に振り分けて送る」
	b.mu.Lock()
	defer b.mu.Unlock()
	for s := range b.subscribers {
		// time.Sleep(time.Second * 10)
		// s.ch <- e
		select {
		case s.ch <- e:
			fmt.Println("受信しました")
		default:
			fmt.Println("timeout")
		}

	}
}

func (b *Broker) closeSubscribers() {
	// 全てのチャンネルをcloseする
	b.mu.Lock()
	defer b.mu.Unlock()
	for s := range b.subscribers {
		close(s.ch)
	}
}

func (b *Broker) Subscribe() *Subscriber {
	b.mu.Lock()
	defer b.mu.Unlock()
	s := &Subscriber{ch: make(chan Event, 10)}
	b.subscribers[s] = struct{}{}
	return s
}

func (b *Broker) Unsubscribe(s *Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.subscribers, s)
	close(s.ch)
}

func (b *Broker) Publish(e Event) {
	// チャンネルに通知を送る
	b.ch <- e
}

// 全てのチャンネルをcloseする
func (b *Broker) Close() {
	close(b.ch)
	<-b.done
}
