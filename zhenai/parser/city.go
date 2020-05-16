package parser

import (
	"github.com/crawler/crawler/engine"
	"regexp"
)
//<div class="userbox" data-uid="3363709">
const cityRe  =`<div class="userbox" data-uid="([0-9]+)">`

func ParseCity(contents []byte,_ string) engine.ParseResult{
	re:=regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result:=engine.ParseResult{}
	for _,m:=range matches{
		url:="http://www.7799520.com/user/"+string(m[1])+".html"
		result.Requests=append(
			result.Requests,engine.Request{
				Url:url,
				Parser: NewProfileParser(string(m[1])),
			})
	}
	return  result
}

