package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/persist/client"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
)

func main(){
	itemChan,err:=client.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err!=nil{
		panic(err)//起不来就挂掉，用panic
	}
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:  "http://city.7799520.com",
		Parser: engine.NewFuncParser(parser.ParseCityList,"ParseCityList"),
	})
}

//docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2