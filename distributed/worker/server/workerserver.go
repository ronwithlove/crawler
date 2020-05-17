package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/distributed/worker"
	"log"
)

//worker的服务器，CrawlService这里的爬虫服务就是单机版的worker
func main()  {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),worker.CrawlService{}))
}
