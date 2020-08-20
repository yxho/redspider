package parse

import (
	"redspider/engine"
	"regexp"
)

const BooklistRe = `<a href="([^"]+" title="([^"}+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(BooklistRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _,m:=range matches{
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
}
