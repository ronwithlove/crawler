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
	limit:=3
	for _,m:=range matches{
		result.Items=append(result.Items,"City "+string(m[2]))//城市名
		result.Requests=append(
			result.Requests,engine.Request{
			Url:	string(m[1]),//url
			ParserFunc: ParseCity,
			})
		limit--//就加载10个城市，要不然太多，看个结果要等半天
		if limit==0{
			break
		}
	}
	return  result
}

