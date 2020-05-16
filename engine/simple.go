package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request){//...表示传入任意个数的Request
	var requests []Request
	for _,r:= range seeds{
		requests=append(requests,r)
	}//先把传入的Request组装好

	for len(requests)>0{
		r:=requests[0]//把第一个request拿出来
		requests=requests[1:]

		parseResult,err:= Worker(r)
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

