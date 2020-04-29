package engine

import (
	"log"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)//interface里的函数不需要名字, Submit是方法，Request是传入参数
	ConfigureMasterWorkerChan(chan Request)//这里也是,不需要方法名，Requestl类型的 channel
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	in:=make (chan Request)
	out:=make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)//通过这个方法把in放到scheduler中去
	//go程并发
	for i:=0; i<e.WorkerCount; i++{
		createWorker(in,out)
	}

	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount:=0
	for{
		result:=<-out//work会输出到out,这里去收
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