package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} //用interface表示類型都可傳入
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
