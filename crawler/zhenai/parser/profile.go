package parser

import (
	"fmt"
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/model"
	"regexp"
	"strconv"
)

/**
解析个人页面数据，获取年龄，是否结婚等数据
*/
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+婚)</div>`)
var idUrlRe = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)
var guessRe = regexp.MustCompile(``)


func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extraceString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extraceString(contents, marriageRe)

	//不再深度遍历解析了，只返回数据
	result := engine.ParseResult{
		Items: []engine.Item{{
			Url:     url,
			Type:    "zhenai",
			Id:      extraceString([]byte(url), idUrlRe),
			Payload: profile,
		}},
	}
	fmt.Printf("profile:%v", profile)

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}
	return result
}

func extraceString(contents []byte, re *regexp.Regexp) string {
	//年龄等找第一个就可以了
	match := marriageRe.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile", p.userName
}

func NewProfileParser(name string) *ProfileParser{
	return &ProfileParser{userName:name}
}