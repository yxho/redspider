package client

import (
	"redspider/distributer/rpcManager"
	"redspider/distributer/worker"
	"redspider/engine"
)

func CreateProcessor()(engine.Processor,error){
	client,err:=rpcManager.NewClient(":1235")
	if err!=nil{
		return nil,err
	}

	return func(req engine.Request)(engine.ParseResult,error){
		sReq:=worker.SerializeRequest(req)
		var sResult worker.ParseResult

		err:=client.Call("CrawlService.Process",sReq,&sResult)
		if err!=nil{
			return engine.ParseResult{},nil
		}
		return worker.DeserializeResult(sResult),nil
	},nil
}
