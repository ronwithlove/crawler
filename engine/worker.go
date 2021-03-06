package engine

import (
	"github.com/crawler/crawler/fetcher"
	"log"
)

//Worker 把fetcher和parser提取出来
func Worker(r Request) (ParseResult,error){
	//	log.Printf("Fetching %s", r.Url)
	body, err:=fetcher.Fetch(r.Url)//fetcher
	if err!=nil{//如果出错，打印下，不可以return，
		log.Printf("Fetcher:error fetching url %s: %v",r.Url,err)
		return  ParseResult{},err//这里ParserResult是个结构，不是指针，不能return nil
	}

	return r.Parser.Parse(body,r.Url),nil//传入content和Url
}
