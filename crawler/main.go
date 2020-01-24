package main

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/persist"
	"learnCrawler/crawler/scheduler"
	"learnCrawler/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemServer("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(
		engine.Request{
			Url:    "http://www.zhenai.com/zhenghun",
			Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
		})

	//e.Run(engine.Requests{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
