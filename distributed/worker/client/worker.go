package client

import (
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/worker"
	"github.com/crawler/crawler/engine"
	"net/rpc"
)


//爬的部分，通过调用rpc，让远程的这个方法去爬
func CreateProcessor(clientChan chan *rpc.Client)engine.Processor{//这里删除自己建客户端，改成从外面传进来
	//这里使用函数是编程，写了一个和engine.RequestProcessor同样签名的匿名函数
	return func(req engine.Request) (engine.ParseResult,  error) {
		sReq:=worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c:=<-clientChan
		err:=c.Call(config.CrawlServiceRpc,sReq,&sResult)
		if err!=nil{
			return  engine.ParseResult{},err
		}
		return worker.DeserializeResult(sResult),nil
	}
}