package parser

import (
	"github.com/crawler/crawler/engine"
	"regexp"
)

//const cityListRe=`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const cityListRe=`<a href="(http://city.7799520.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	re:=regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result:=engine.ParseResult{}

	for k,m:=range matches{
		result.Items=append(result.Items,"City "+string(m[2]))//城市名
		result.Requests=append(
			result.Requests,engine.Request{
			Url:	string(m[1]),//url
			ParserFunc: ParseCity,
			})
		if k==2{//就找前3个城市，每个城市有15个会员
			break
		}
	}
	return  result
}

