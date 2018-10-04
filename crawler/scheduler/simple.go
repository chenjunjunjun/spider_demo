// 并发版


package scheduler

import "spider_demo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}



func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func() {
		s.workerChan <- r
	}()

}





