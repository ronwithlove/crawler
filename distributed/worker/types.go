package worker

import (
	"errors"
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/zhenai/parser"
	"log"
)

//这里需要序列化自己的Parser转换成engine的，才可以在网络上传播 ，	serializ: 自己=>别人(可网上传播）
//同时也无法直接用engine的，需要通过反序列化转化成自己的，才可以用	deserialize:别人的(网上传播的)=>自己

//序列化Parser
type SerializedParser struct{
	Name string//方法名 ,如：ParseCityList，ProfileParser
	Args interface{}//这里ParseCityList对应的就是nil,ProfileParser对应的就是传入username
}
//{"ParseCityList",nil},{"ProfileParser",username}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct{
	Requests []Request //自己的Request,里面包含的Parser是可以在网络上传的，
	// engine.types.go中的Parser是没法在网络上传的，因为里面的Parser个interfadce
	Items []engine.Item
}

//把engine的Request转换成自己的Request
func SerializeRequest(r engine.Request) Request{
	name,args:= r.Parser.Serialize()
	return Request{
		Url:    r.Url,
		Parser: SerializedParser{
			Name:name,
			Args:args,
		},
	}
}

//把engine的ParseResult转换成自己的ParseResult
func SerializeResult(r engine.ParseResult) ParseResult{
	result:=ParseResult{
		Items:r.Items,
	}
	for _,req:=range r.Requests{
		result.Requests =append(result.Requests,SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request,error){
	parser, err := deserializeParser(r.Parser)
	if err!=nil{
		return engine.Request{},err
	}
	return  engine.Request{
		Url:    r.Url,
		Parser: parser,
	},nil
}

func DeserializeResult(r ParseResult)engine.ParseResult{
	result:=engine.ParseResult{
		Items:r.Items,
	}
	for _,req:=range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err!=nil{
			log.Printf("error deserializing request: %v", err)
			continue//如果有错的req，提示，然后忽略，继续下一个request
		}
		result.Requests =append(result.Requests,engineReq)
	}
	return result
}

//把自己的parser 转成engine的
func deserializeParser(p SerializedParser) (engine.Parser,error){
	switch p.Name{
	case config.ParseCityList:
		return  engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),nil
	case config.ParseCity:
		return  engine.NewFuncParser(parser.ParseCity,config.ParseCity),nil
	case config.NilParser:
		return  engine.NilParser{},nil
	case config.ParseProfile:
		if userId,ok:=p.Args.(string);ok{
			return  parser.NewProfileParser(userId),nil
		}else{
			return nil,fmt.Errorf("invalid arg: %v",p.Args)
		}
	default:
		return nil, errors.New("unknown parser name:"+p.Name)
	}

}

