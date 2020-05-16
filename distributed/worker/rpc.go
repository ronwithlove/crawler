package worker

import "github.com/crawler/crawler/engine"

//这里放rpc服务
type CrawlService struct {

}

//爬虫方法，就是一个rpc的work，但是要转化一下里面的参数，才可以在网上传播
func (CrawlService) Process (req Request, result *ParseResult)error{
	engineReq, err := DeserializeRequest(req)//先把自己的request转成engine的request
	if err!=nil{
		return err
	}

	engineResult,err:=engine.Worker(engineReq)//通过worker得到engine的result
	if err!=nil{
		return err
	}
	//再转成自己的result,然后直接赋值给自己的result
	*result = SerializeResult(engineResult)
	return nil
}

