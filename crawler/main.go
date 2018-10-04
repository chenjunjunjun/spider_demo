package main

import (
	"spider_demo/crawler/engine"
	"spider_demo/crawler/persist"
	"spider_demo/crawler/scheduler"
	"spider_demo/crawler/zhenai/parser"
)

func main()  {

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}


	e.Run(engine.Request{
		Url: 	"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

