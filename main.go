package main

import (
	"redspider/engine"
	"redspider/parse"
	"redspider/persist"
	"redspider/scheduler"
)

func main() {
	itemsave, err := persist.ItemSave()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
