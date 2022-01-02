package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
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
