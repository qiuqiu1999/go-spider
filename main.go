package main

import (
	"go-spider/engine"
	"go-spider/scheduler"
	"go-spider/zhenai/parser"
)

func main() {
	r := engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	/*r := engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/wuhan/nv",
		ParserFunc: parser.ParseCity,
	}*/
	engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerNum: 50,
	}.Run(r)

}
