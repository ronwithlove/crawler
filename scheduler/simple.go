package scheduler

import (
	"github.com/crawler/crawler/engine"
)

type SimpleScheduler struct{
	workerChan chan engine.Request
}

//内容被改变就要用指针 *
func (s *SimpleScheduler)ConfigureMasterWorkerChan(c chan engine.Request){
	s.workerChan=c
}

func (s *SimpleScheduler) Submit(r engine.Request)  {
	//send request to worker chan
	//这里要用goroutine把in尽快收走
	go func() {s.workerChan<-r}()
}
