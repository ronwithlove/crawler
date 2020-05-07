package engine

type Request struct{
	Url string
	ParserFunc func([]byte) ParseResult
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

func NilParser([]byte) ParseResult{
	return ParseResult{}
}