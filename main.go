package main

import (
	"fmt"
	"go-notification/pubsub"
	"sync"
)

func main() {
	br := pubsub.NewBrocker()

	sub1 := br.Subscribe()
	sub2 := br.Subscribe()
	br.Start()

	var wg sync.WaitGroup

	wg.Go(func() {
		for m1 := range sub1.Events() {
			fmt.Println("m1の内容")
			fmt.Printf("%s\n", m1.ID)
			fmt.Printf("%s\n", m1.Text)
		}
	}) // メッセージを受信する

	wg.Go(func() {
		for m2 := range sub2.Events() {
			fmt.Println("m2の内容")
			fmt.Printf("%s\n", m2.ID)
			fmt.Printf("%s\n", m2.Text)
		}
	})

	// メッセージを送る
	br.Publish(pubsub.Event{ID: "100", Text: "a"})
	br.Publish(pubsub.Event{ID: "200", Text: "b"})
	br.Publish(pubsub.Event{ID: "300", Text: "c"})

	br.Close()

	wg.Wait()

}
