package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"redspider/engine"
)

func ItemSave() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemcount := 0
		for {
			item := <-out
			log.Printf("Item saver :Got$%d,%v", itemcount, item)
			Save(client, item)
			itemcount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	_, err := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item).Do(context.Background())

	if err != nil {
		panic(err)
	}

	return nil
}
