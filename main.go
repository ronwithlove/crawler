package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
)

func main(){
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},//换这个也行
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		//Url:  "http://www.zhenai.com/zhenghun",
		Url:  "http://city.7799520.com",
		ParserFunc: parser.ParseCityList,
	})
}
