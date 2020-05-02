package scheduler

import (
	"github.com/crawler/crawler/engine"
)

type SimpleScheduler struct{
	workerChan chan engine.Request
}

//内容被改变就要用指针 *

func (s *SimpleScheduler) Submit(r engine.Request)  {
	//send request to worker chan
	//这里要用goroutine把in尽快收走,这里就等于in<-seed
	go func() {s.workerChan<-r}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan=make(chan engine.Request)
}
