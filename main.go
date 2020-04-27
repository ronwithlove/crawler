package main

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/zhenai/parser"
)

func main(){
	engine.Run(engine.Request{
		//Url:  "http://www.zhenai.com/zhenghun",
		Url:  "http://city.7799520.com",
		ParserFunc: parser.ParseCityList,
	})
}
