package main

import (
	"redspider/distributer/client"
	client2 "redspider/distributer/worker/client"
	"redspider/engine"
	"redspider/parse"
	"redspider/scheduler"
)

func main() {
	itemsave, err := client.ItemSave(":1234")

	process, err := client2.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
		process,
	}
	e.Run(engine.Request{
		Url:   "https://book.douban.com/",
		Parse: engine.NewFuncParse(parse.ParseTag, "booklist"),
	})
}
