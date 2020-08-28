package parse

import (
	"fmt"
	"redspider/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseTag(content []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range match {
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))

		//result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:   "https://book.douban.com" + string(m[1]),
			Parse: engine.NewFuncParse(ParseBookList, "booklist"),
		})
	}
	return result
}
