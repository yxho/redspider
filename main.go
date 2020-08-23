package main

import (
	"redspider/engine"
	"redspider/parse"
	"redspider/scheduler"
)

func main() {
	e:=engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
	}
	e.Run(engine.Request{
		Url: "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}


