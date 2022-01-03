package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerNum int
	ItemChan  chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkReady(w chan Request)
}

func (c ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerNum; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, seed := range seeds {
		fmt.Printf("add request: %s\n", seed.Url)
		c.Scheduler.Submit(seed)
	}
	for {
		result := <-out

		for _, item := range result.Item {
			//fmt.Printf("Got item %+v \n", item)
			item := item
			go func() { c.ItemChan <- item }()
		}

		for _, v := range result.Request {
			c.Scheduler.Submit(Request{
				Url:        v.Url,
				ParserFunc: v.ParserFunc,
			})
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, r ReadyNotifier) {
	go func() {
		for {
			r.WorkReady(in)
			req := <-in
			result, err := work(req)
			if err != nil {
				fmt.Printf("url: %v , error:%s\n", req.Url, err)
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

//去重
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
