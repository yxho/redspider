package scheduler

import "redspider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) configureWorkChan(c chan engine.Request) {
	s.workerChan = c
}
