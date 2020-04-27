package parser

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"regexp"
)

//预先编译下,而不是放到程序运行的时候在去做这步，提高效率
//`<span class="marrystatus">未婚</span>`
var nameRe=regexp.MustCompile(`<span class="nick c3e">([^<]+)</span>`)
var marrigageRe=regexp.MustCompile(`<span class="marrystatus">([^<]+)</span>`)
//`<span class="age s1">23岁</span>`
var ageRe=regexp.MustCompile(`<span class="age s1">([\d]+)岁</span>`)//\d就是1个到多个数字,加上括号，提取数字

func ParseProfile(contents []byte,userid string) engine.ParseResult{
	profile:=model.Profile{}

	//Userid
	profile.UserId=userid
	profile.Name=extractString(contents,nameRe)
	profile.Age=extractString(contents,ageRe)
	profile.Marriage=extractString(contents,marrigageRe)

	result:=engine.ParseResult{
		Items:[]interface{}{profile},//这里的items放的是profile的结构体
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	match:=re.FindSubmatch(contents)//这里不用findall了，这里就一个
	if len(match)>=2{//一般至少有2个，一个是提取的一窜string，一个是他的submatch
		return string(match[1])
	}else{
		return ""
	}
}

