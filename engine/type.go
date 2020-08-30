package engine

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}
type Item struct {
	Url     string
	Type    string
	Payload interface{}
}

type Request struct {
	Url   string
	Parse Parser
}

type NilParse struct {
}

func (n NilParse) Parser(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParse) Serialize() (name string, args interface{}) {
	return "Nilparse", nil
}

type ParseFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	Parser_ ParseFunc
	name    string
}

func (f FuncParser) Parser(contents []byte, url string) ParseResult {
	return f.Parser_(contents, url)
}

func (f FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParse(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		Parser_: p,
		name:    name,
	}
}
