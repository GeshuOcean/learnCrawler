package main

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})
}
