package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	inc int
	mu  sync.Mutex
}

func (c *Counter) Run(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		// go increment()
		wg.Go(c.increment)
	}
	wg.Wait()
}

func (c *Counter) increment() {
	c.mu.Lock()
	c.inc++
	c.mu.Unlock()
}
func main() {

	counter := Counter{}
	counter.Run(1000)

	fmt.Printf("%d", counter.inc)

}
