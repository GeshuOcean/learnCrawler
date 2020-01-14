package engine

import (
	"learnCrawler/crawler/fetcher"
	"log"
)


type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult,err := e.worker(r)
		if err != nil {
			continue
		}

		//解析出来的新请求添加到队列，并打印爬取的数据
		requests = append(requests, parseResult.Request...)
		for _, item := range parseResult.Items {
			log.Printf("Goe iten %s", item)
		}
	}
}

//请求url拉取数据	Fetch比较耗时
func (e SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fecher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	//调用Request中方法解析url请求结果中所需数据
	return r.ParserFunc(body),nil
}
