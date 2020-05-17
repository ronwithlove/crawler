package client

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/distributed/worker"
	"github.com/crawler/crawler/engine"
)

//爬的部分，通过调用rpc，让远程的这个方法去爬
func CreateProcessor() (engine.Processor,error){
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d",config.WorkerPort0))
	if err != nil {
		return  nil,err
	}

	//这里使用函数是编程，写了一个和engine.RequestProcessor同样签名的匿名函数
	return func(req engine.Request) (engine.ParseResult,  error) {
		sReq:=worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err:=client.Call(config.CrawlServiceRpc,sReq,&sResult)
		if err!=nil{
			return  engine.ParseResult{},err
		}
		return worker.DeserializeResult(sResult),nil
	},nil
}