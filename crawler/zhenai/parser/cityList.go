package parser

import (
	"fmt"
	"learnCrawler/crawler/engine"
	"regexp"
)

const cityListRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches {

		result.Items = append(result.Items, "City "+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		fmt.Printf("City: %s ,URL: %s ", m[2], m[1])
		limit--
		if limit <= 0 {
			break
		}
		fmt.Println()
	}
	return result
}
