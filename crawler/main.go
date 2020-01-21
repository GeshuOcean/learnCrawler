package main

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/persist"
	"learnCrawler/crawler/scheduler"
	"learnCrawler/crawler/zhenai/parser"
)

func main() {
	itemChan,err:=persist.ItemServer("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:   itemChan ,
	}

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
