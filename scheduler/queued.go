package scheduler

import (
	"go-spider/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQueue) > 0 && len(workQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWork = workQueue[0]
			}
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChan:
				workQueue = append(workQueue, w)
			case activeWork <- activeRequest:
				requestQueue = requestQueue[1:]
				workQueue = workQueue[1:]
			}
		}
	}()
}
