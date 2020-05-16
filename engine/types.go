package engine

type ParserFunc func(contents []byte,url string) ParseResult

type Request struct{
	Url string
	Parser Parser //这里变成接口了
}

type Parser interface {
	Parse(contents []byte,url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseResult struct {
	Requests []Request
	Items 	[]Item
}

type Item struct{
	Url string
	Id string
	Payload interface{}//可以是任意的，这里用model.Proflie
}


type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return  ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}

//不管他struct内有啥参数，只要继承了Parser interface，他就可以当Parser来用
type FuncParser struct {
	parser ParserFunc //放parser方法
	name string	//这个方法的名字
}

//工厂函数来建FuncParser
func NewFuncParser(p ParserFunc, name string) *FuncParser{
	return &FuncParser{
		parser:p,
		name:name,
	}
}

//继承了Parser interface,city, cityList公用的Parse，不需要另外传入参数
func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name,nil
}
