package main

import (
	"fmt"
	"go-notification/publisher"
	"go-notification/pubsub"
	"go-notification/subscriber"
)

func main() {
	// br := pubsub.NewBrocker()

	// sub1 := br.Subscribe()
	// sub2 := br.Subscribe()
	// br.Start()
	// br.Unsubscribe(sub1)

	// var wg sync.WaitGroup

	// wg.Go(func() {
	// 	for m1 := range sub1.Events() {
	// 		fmt.Println("m1の内容")
	// 		fmt.Printf("%s\n", m1.ID)
	// 		fmt.Printf("%s\n", m1.Text)
	// 	}
	// }) // メッセージを受信する

	// wg.Go(func() {
	// 	for m2 := range sub2.Events() {
	// 		fmt.Println("m2の内容")
	// 		fmt.Printf("%s\n", m2.ID)
	// 		fmt.Printf("%s\n", m2.Text)
	// 	}
	// })

	// メッセージを送る
	// br.Publish(pubsub.Event{ID: "100", Text: "a"})
	// br.Publish(pubsub.Event{ID: "200", Text: "b"})
	// br.Publish(pubsub.Event{ID: "300", Text: "c"})

	// レースコンディジョンを起こす
	// go br.Publish(pubsub.Event{ID: "300", Text: "c"})
	// go br.Subscribe()

	// br.Close()

	// wg.Wait()

	// ブローカー起動
	br := pubsub.NewBrocker()
	br.Start()

	// サブスクライバー起動
	subscriber := subscriber.NewService(br, func(e pubsub.Event) {
		fmt.Println("eventの内容")
		fmt.Printf("%s\n", e.ID)
		fmt.Printf("%s\n", e.Text)
	})
	subscriber.Start()

	// パブリッシャー起動
	pub := publisher.NewPublisher(br)
	pub.Publish(pubsub.Event{ID: "100", Text: "a"})
	br.Close()
	subscriber.Wait()
}
