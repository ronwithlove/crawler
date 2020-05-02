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
	//先创建好这些work，workerChan channel没的时候就等着呗
	for i:=0; i<e.WorkerCount; i++{
		createWorker(in,out)
	}
	//开始往workerChan写入request了，然后就会自动运行把结果输出到out
	//这一步就相当于把in 带入到createWorker里
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount:=0
	for{//一直循环读取out
		result:=<-out//这里去收out
		for _, item:=range result.Items{//把结果打印出来
			log.Printf("Got item #%d: %v",itemCount,item)
			itemCount++
		}
		//把第二个属性request	再放到workChian，再去让他执行
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