package engine

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)//interface里的函数不需要名字, Submit是方法，Request是传入参数
	WorkerChan()chan Request
	Run()
}

type ReadyNotifier interface {//这是一个结构体
	WorkerReady(chan Request)//结构体里面有一个这个方法
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	out:=make(chan ParseResult)
	e.Scheduler.Run()//request分配给worker，一个并行无限循环的方法
	//先创建好这些work，在没有workerChan channel的时候就等着
	for i:=0; i<e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r:=range seeds{
		e.Scheduler.Submit(r)//把request 都给到requestChan
	}

	for{//一直循环读取out
		result:=<-out//这里去收out,如果有就收走
		for _, item:=range result.Items{//把结果打印出来
			go func(){e.ItemChan<-item}()//不用go程会死锁
		}
		//把第二个属性request	再放到requestChan，再去让去等着，等有空闲的 request chan chan就又可以被执行了
		for _, request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(in chan Request, out chan ParseResult,ready ReadyNotifier) {
	go func() {
		for{//这里使用了ReadyNotifier.WorkerReady方法把reuest channel传给了 workerChan
			ready.WorkerReady(in)//如果有闲置输入，就自动去workerQ等着了
			request:=<-in//直到in在scheduler.Run接收到了request，这里的request才读的出来，
			result,err:=worker(request)//于是进入这行，开工，返回一个ParseResult
			if err!=nil{
				continue
			}
			out<-result//out读取ParseResult，这里不需要返回，out就是外面那个out
		}
	}()
}