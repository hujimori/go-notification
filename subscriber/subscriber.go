package subscriber

import (
	"go-notification/pubsub"
	"sync"
)

var wg sync.WaitGroup

type Service struct {
	br      *pubsub.Broker
	handler func(pubsub.Event)
	wg      sync.WaitGroup
}

func NewService(br *pubsub.Broker, handler func(pubsub.Event)) *Service {
	return &Service{
		br:      br,
		handler: handler,
	}
}

func (s *Service) Start() {
	s.wg.Go(func() {
		for e := range s.br.Subscribe().Events() {
			s.handler(e)
		}
	})
}
func (s *Service) Wait() {
	s.wg.Wait()
}

func (s *Service) Stop() {

}
