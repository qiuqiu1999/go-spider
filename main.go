package main

import (
	"github.com/garyburd/redigo/redis"
	"go-spider/engine"
	"go-spider/persist"
	"go-spider/scheduler"
	"go-spider/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	redisConn, err := redis.Dial("tcp", "192.168.5.170:6379")
	if err != nil {
		panic(err)
	}
	defer redisConn.Close()

	engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerNum: 50,
		ItemChan:  itemChan,
		RedisConn: redisConn,
	}.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
