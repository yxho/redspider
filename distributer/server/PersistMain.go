package main

import (
	"gopkg.in/olivere/elastic.v5"
	"redspider/distributer/persist"
	"redspider/distributer/rpcManager"
)

func main() {
	serveRpc(":1234")
}
func serveRpc(host string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcManager.ServeRpc(host, &persist.ItemService{
		Client: client,
	})
}
