package engine

import (
	"log"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)//这里的Request是interface{}类型，不需要名字, Submit是方法，Request是传入参数
	ConfigureMasterWorkerChan(chan Request)
}
func (e *ConcurrentEngine) Run(seeds ...Request){
	in:=make (chan Request)
	out:=make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}

	//go程并发，无限循环
	for i:=0; i<e.WorkerCount; i++{
		createWorker(in,out)
	}

	itemCount:=0
	for{//无限循环
		result:=<-out
		for _, item:=range result.Items{
			log.Printf("Got item #%d: %v",itemCount,item)
			itemCount++
		}

		for _, request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for{
			request:=<-in
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			out<-result
		}
	}()
}