package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request [] Request	//新的请求
	Items   []interface{} //解析的数据
}

func Nilparser([]byte) ParseResult {
	return ParseResult{}
}
