package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemcount := 0
		for {
			item := <-out
			log.Printf("Item saver :Got$%d,%v", itemcount, item)
			save(item)
			itemcount++
		}
	}()
	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	_, err = client.Index().Index("dating_profile").Type("douban").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}
}
