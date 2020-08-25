package main

import (
	"redspider/engine"
	"redspider/parse"
	"redspider/persist"
	"redspider/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		persist.ItemSave(),
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
