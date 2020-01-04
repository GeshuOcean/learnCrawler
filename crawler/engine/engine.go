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

		requests = append(requests, parseResult.Request...)
		for _, item := range parseResult.Items {
			log.Printf("Goe iten %s", item)
		}

	}
}

//Fetch比较耗时
func (e SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fecher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body),nil
}
