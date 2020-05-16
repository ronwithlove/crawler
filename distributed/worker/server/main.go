package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/distributed/worker"
	"log"
)

//服务器的main
func main()  {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),worker.CrawlService{}))

}
