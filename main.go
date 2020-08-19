package main

import (
	"redspider/engine"
	"redspider/parse"
)

func main() {
	engine.Run(engine.Request{
		Url: "https://book.douban.com",
		ParseFunc: parse.ParseContent,
	})
}


