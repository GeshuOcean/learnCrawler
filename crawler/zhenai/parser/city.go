package parser

import (
	"fmt"
	"learnCrawler/crawler/engine"
	"regexp"
)

/**
解析城市，获取用户个人信息url
*/
var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`(href="http://www.zhenai.com/zhenghun/shanhai/[^"]+)"`)
)

func ParseCity(contents []byte,url string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches {
		//result.Items = append(result.Items, "User "+name)
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
		limit--
		if limit <= 0 {
			break
		}
		fmt.Println()
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
