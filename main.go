package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/persist"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
)

func main(){
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:	persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:  "http://city.7799520.com",
		ParserFunc: parser.ParseCityList,
	})
}
