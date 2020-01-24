package parser

import (
	"learnCrawler/crawler/engine"
	"regexp"
)

/**
解析获取城市列表
*/
//正则中括号（）中为需要提取的数据<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f>阿坝</a>
const cityListRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//城市列表解析出城市数据交给城市解析器
func ParseCityList(contents []byte, url string) engine.ParseResult {
	//正则
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		//新request添加到本次解析要返回结果中的Request队列中
		//result.Items = append(result.Items, engine.Item{
		//	Url:     "",
		//	Type:    "",
		//	Id:      "",
		//	Payload: nil,
		//})
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		limit--
		if limit <= 0 {
			break
		}
	}
	return result
}
