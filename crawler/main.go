package main

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/persist"
	"learnCrawler/crawler/scheduler"
	"learnCrawler/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(),
	}

	//e.Run(
	//	engine.Request{
	//		Url:        "http://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
