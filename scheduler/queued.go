package scheduler

import "github.com/crawler/crawler/engine"

type QueuedScheduler struct{
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

//request channel, request好了就把request送进去
func (s *QueuedScheduler)Submit(r engine.Request){
	s.requestChan<-r
}
//woker channel channle, worker channel 好了就把worker channel送进去
func (s *QueuedScheduler)WorkerReady(w chan engine.Request){
	s.workerChan<-w
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return  make(chan engine.Request)
}


func (s *QueuedScheduler) Run(){//这里要加*，就是指针，因为使用了s.workerChan这个变量，改变了他的内容
	s.workerChan=make(chan chan engine.Request)
	s.requestChan=make(chan engine.Request)
	go func() {//并行
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{//无限循环
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0&&len(workerQ)>0{//这个时候requestChan和workerChan队列同时都有空闲的
				activeRequest=requestQ[0]
				activeWorker=workerQ[0]
			}
			select {
			case r:=<-s.requestChan://request 有就接着
				requestQ=append(requestQ,r)//放到队列等着
			case w:=<-s.workerChan://有空的request chan 就接着
				workerQ=append(workerQ,w)//放到管道中
			case activeWorker<-activeRequest://只有上面两个列队都有东西了，这行代码才会被执行
				requestQ=requestQ[1:]//然后把他从序列里移除
				workerQ=workerQ[1:]
			}
		}

	}()
}
