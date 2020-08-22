package main

import (
	"redspider/engine"
	"redspider/parse"
)

func main() {
	engine.Run(engine.Request{
		Url: "https://book.douban.com/subject/30293801",
		ParseFunc: parse.ParseBookDetail,
	})
}


