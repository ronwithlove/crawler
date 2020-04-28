package engine

import (
	"github.com/crawler/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){//...表示传入任意个数的Request
	var requests []Request
	for _,r:= range seeds{
		requests=append(requests,r)
	}//先把传入的Request组装好

	for len(requests)>0{
		r:=requests[0]//把第一个request拿出来
		requests=requests[1:]

		parseResult,err:=worker(r)
		if err!=nil{
			continue
		}

		//分析出两样东西，1.又得到了一个子的request，加进去
		requests=append(requests,parseResult.Requests...)//把parseResult.Requests里面所有元素一个个加进去
		//2.得到items，打印出来
		for _,item:=range parseResult.Items{
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult,error){
	log.Printf("Fetching %s", r.Url)
	//1.通过网页抓取内容
	body, err:=fetcher.Fetch(r.Url)
	if err!=nil{//如果出错，打印下，不可以return，
		// 因为有很多个request，不可能一个不对，全挂了，如果这样的话爬虫就没法工作了
		log.Printf("Fetcher:error fetching url %s: %v",
			r.Url,err)
		return  ParseResult{},err//这里ParserResult是个结构，不是指针，不能return nil
	}

	//2.使用request里的分析方法来分析抓取的内容
	return r.ParserFunc(body),nil
}
