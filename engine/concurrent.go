package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerNum int
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
	itemCount := 0
	for {
		result := <-out
		for _, v := range result.Item {
			itemCount++
			fmt.Printf("id:%d Got item %v \n", itemCount, v)
		}

		for _, v := range result.Request {
			//fmt.Printf("id:%d  add request: %s \n", itemCount, v.Url)
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
