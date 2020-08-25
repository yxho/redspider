package engine

import (
	"fmt"
	"log"
	"redspider/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}
type Scheduler interface {
	Submit(Request)
	//configureWorkChan()
	Run()
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		CreateWork(out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got item:%d,%v", itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWork(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(request Request) (ParseResult, error) {
	fmt.Printf("Fetch url:%s\n", request.Url)
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch Error:%s\n", request.Url)
		return ParseResult{}, err
	}
	return request.ParseFunc(body), nil

}
