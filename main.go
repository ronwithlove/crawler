package main

import (
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/persist"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
)

func main(){
	itemChan,err:=persist.ItemSaver("dating_profile")
	if err!=nil{
		panic(err)//起不来就挂掉，用panic
	}
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor:engine.Worker,//单机版直接调Worker
	}
	e.Run(engine.Request{
		Url:  "http://city.7799520.com",
		Parser: engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),
	})
}
