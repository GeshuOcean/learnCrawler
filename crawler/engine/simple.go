package engine

import (
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

		parseResult,err := worker(r)
		if err != nil {
			continue
		}

		//解析出来的新请求添加到任务队列，并打印爬取的数据
		requests = append(requests, parseResult.Request...)
		for _, item := range parseResult.Items {
			log.Printf("Goe iten %s", item)
		}
	}
}




