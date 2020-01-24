package engine

import (
	"learnCrawler/crawler/fetcher"
	"log"
)

//请求url拉取数据	Fetch比较耗时
func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fecher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	//调用Request中方法解析url请求结果中所需数据
	return r.Parser.Parse(body, r.Url), nil
}
