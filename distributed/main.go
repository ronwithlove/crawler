package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	itemsaver "github.com/crawler/crawler/distributed/persist/client"
	worker "github.com/crawler/crawler/distributed/worker/client"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
)

//分布式爬虫的main
func main(){
	//启一个itemsaver的客户端
	itemChan,err:=itemsaver.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err!=nil{
		panic(err)//起不来就挂掉，用panic
	}

	processor,err:=worker.CreateProcessor()
	if err!=nil{
		panic(err)
	}
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor:processor,
	}
	e.Run(engine.Request{
		Url:  "http://city.7799520.com",
		Parser: engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),
	})
}

//docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2