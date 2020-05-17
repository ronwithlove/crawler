package parser

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"regexp"
)

//预先编译下,而不是放到程序运行的时候在去做这步，提高效率
//`<span class="marrystatus">未婚</span>`
var nameRe = regexp.MustCompile(`<span class="nick c3e">([^<]+)</span>`)
var marrigageRe = regexp.MustCompile(`<span class="marrystatus">([^<]+)</span>`)

//`<span class="age s1">23岁</span>`
var ageRe = regexp.MustCompile(`<span class="age s1">([\d]+)岁</span>`) //\d就是1个到多个数字,加上括号，提取数字

func parseProfile(contents []byte, url string, userid string) engine.ParseResult {
	profile := model.Profile{}

	//Userid
	profile.UserId = userid
	profile.Name = extractString(contents, nameRe)
	profile.Age = extractString(contents, ageRe)
	profile.Marriage = extractString(contents, marrigageRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      userid,
				Payload: profile,
			},
		}, //这里的items放的是profile的结构体
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents) //这里不用findall了，这里就一个
	if len(match) >= 2 {               //一般至少有2个，一个是提取的一窜string，一个是他的submatch
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userID string
}

//Profile的Parser相比其他的还需要从外部传入一个userID
func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents,url,p.userID)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile",p.userID//序列化的时候也要这个参数
	//return "ProfileParser",p.userID//序列化的时候也要这个参数
}

//只要他return的是继承了Parser就可以了
func NewProfileParser(userid string) *ProfileParser{
	return &ProfileParser{//返回的是指针，所以这里要用&，地址
		userID:userid,
	}
}
