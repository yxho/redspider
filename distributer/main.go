package main

import (
	"redspider/distributer/client"
	"redspider/engine"
	"redspider/parse"
	"redspider/scheduler"
)

func main() {
	itemsave, err := client.ItemSave(":1234")
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
		Parse: engine.NewFuncParse(parse.ParseTag,"booklist"),
	})
}
