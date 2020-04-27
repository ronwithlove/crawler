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
		result.Items=append(result.Items,"UserID "+string(m[1]))//usrid
		result.Requests=append(
			result.Requests,engine.Request{
				Url:	"http://www.7799520.com/user/"+string(m[1])+".html",//url
				ParserFunc: engine.NilParser,
			})
	}
	return  result
}