package parser

import (
	"github.com/crawler/crawler/engine"
	"regexp"
)
//<div class="userbox" data-uid="3363709">
const cityRe  =`<div class="userbox" data-uid="([0-9]+)">`

func ParseCity(contents []byte) engine.ParseResult{
	re:=regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result:=engine.ParseResult{}
	for _,m:=range matches{
		userid:=string(m[1])
		result.Items=append(result.Items,"UserID "+userid)//usrid
		result.Requests=append(
			result.Requests,engine.Request{
				Url:	"http://www.7799520.com/user/"+string(m[1])+".html",//url
				//ParserFunc: ParseProfile,
				//函数式编程
				ParserFunc: func(con []byte) engine.ParseResult{
					return  ParseProfile(con,userid)
				},
			})
	}
	return  result
}