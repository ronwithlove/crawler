package main

import (
	"flag"
	"fmt"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/distributed/worker"
	"log"
)



var port=flag.Int("port",0,"open port to listen")

//使用如下命令启动workerserver，带端口：
//go run workerserver.go --port=9000

//worker的服务器，CrawlService这里的爬虫服务就是单机版的worker
func main()  {
	flag.Parse()
	if *port==0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{}))
}

