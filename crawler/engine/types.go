package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request //新的请求
	Items   []Item    //解析的数据
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func Nilparser([]byte) ParseResult {
	return ParseResult{}
}
