package main

import (
	"flag"
	"github.com/crawler/crawler/config"
	itemsaver "github.com/crawler/crawler/distributed/persist/client"
	"github.com/crawler/crawler/distributed/rpcsupport"
	worker "github.com/crawler/crawler/distributed/worker/client"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/scheduler"
	"github.com/crawler/crawler/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var(
	itemSaverHost=flag.String("itemsaver_host","","itemsaver_host")
	workerHosts=flag.String("worker_hosts","","worker_hosts (用逗号分开)")
)
//分布式爬虫的main
func main(){
	flag.Parse()
	//启一个itemsaver的客户端
	itemChan,err:=itemsaver.ItemSaver(*itemSaverHost)
	if err!=nil{
		panic(err)//起不来就挂掉，用panic
	}

	//建立client池
	pool:=createClientPool(strings.Split(*workerHosts,","))
	processor :=worker.CreateProcessor(pool)

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

//对host一个一个去连，连接完成后形成一个client pool,然后通过chan来发给work
func createClientPool(hosts []string) chan *rpc.Client{
	//建立clients pool
	var clients []*rpc.Client
	for _,h:=range hosts{
		client,err:=rpcsupport.NewClient(h)
		if err==nil{
			clients=append(clients,client)
			log.Printf("Connected to %s", h)
		}else{
			log.Printf("Error connecting to %s: %v",h,err)
		}
	}

	//channel分发,一般在goroutine里，是go语言的常用写法
	out:=make(chan *rpc.Client)
	go func(){
		for{//不停的发
			for _,client:=range clients{
				out<-client//只要被去走，就往里放
			}
		}
	}()
	return out
}
//docker启用elasticsearch命令：
//docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2
//运行main的命令，注意port前要加冒号
//go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"