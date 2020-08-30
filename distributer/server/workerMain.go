package main

import (
	"log"
	"redspider/distributer/rpcManager"
	serv "redspider/distributer/worker/server"
)

func main(){
	log.Fatal(rpcManager.ServeRpc(":1235",&serv.CrawlService{}))
}
