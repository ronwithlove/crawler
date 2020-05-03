package engine

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)//interface里的函数不需要名字, Submit是方法，Request是传入参数
	WorkerChan()chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	out:=make(chan ParseResult)
	e.Scheduler.Run()
	//先创建好这些work，workerChan channel没的时候就等着呗
	for i:=0; i<e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}
	//开始往workerChan写入request了，然后就会自动运行把结果输出到out
	//这一步就相当于把in 带入到createWorker里
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}

	for{//一直循环读取out
		result:=<-out//这里去收out
		for _, item:=range result.Items{//把结果打印出来
			go func(){e.ItemChan<-item}()
		}
		//把第二个属性request	再放到workChian，再去让他执行
		for _, request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult,ready ReadyNotifier) {
	go func() {
		for{
			ready.WorkerReady(in)
			request:=<-in
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			out<-result
		}
	}()
}