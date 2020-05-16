package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/distributed/worker"
	"testing"
	"time"
)

//测试爬虫rpc
func TestCrawlService(t *testing.T) {
	const host = ":9000"
	//先起一个服务器。会持续监听
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second) //等待个1秒钟让服务起来，再执行client

	//新一个客户端
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//模拟一个实际的测试内容
	req := worker.Request{
		Url: "http://www.7799520.com/user/3376375.html",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "3376375",
		},
	}

	var result worker.ParseResult
	//call爬虫rpc,得到result
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {//没错就打印出来
		fmt.Println(result)
	}
}
