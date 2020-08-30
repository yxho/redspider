package server

import (
	"redspider/distributer/worker"
	"redspider/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(req worker.Request, result *worker.ParseResult) error {
	engineReq, err := worker.DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	* result = worker.SerializeResult(engineResult)
	return nil
}
