package main

import (
	"fmt"
	"go-notification/pubsub"
	"sync"
)

func main() {
	br := pubsub.NewBrocker()

	ch := make([]chan pubsub.Event, 2)

	ch[0] = make(chan pubsub.Event)
	ch[1] = make(chan pubsub.Event)

	br.Subscribe(ch[0])
	br.Subscribe(ch[1])

	var wg sync.WaitGroup
	wg.Add(2)
	// メッセージを受信する
	go func() {
		defer wg.Done()

		for m1 := range ch[0] {
			fmt.Println("m1の内容")
			fmt.Printf("%s\n", m1.ID)
			fmt.Printf("%s\n", m1.Text)
		}
	}()

	go func() {
		defer wg.Done()
		for m2 := range ch[1] {
			fmt.Println("m2の内容")
			fmt.Printf("%s\n", m2.ID)
			fmt.Printf("%s\n", m2.Text)
		}
	}()

	// メッセージを送る
	br.Publish(pubsub.Event{ID: "100", Text: "a"})
	br.Publish(pubsub.Event{ID: "200", Text: "b"})
	br.Publish(pubsub.Event{ID: "300", Text: "c"})

	br.Close()

	wg.Wait()

}
