package persist

import (
	"learnCrawler/crawler/model"
	"testing"
)

func TestItemServer(t *testing.T) {
	expected := model.Profile{
		Age:        34,
		Height:     168,
		Weight:     60,
		Income:     "1K",
		Gender:     "女",
		Name:       "孙宁",
		Xinzuo:     "牧羊座",
		Occupation: "人事",
		Marriage:   "离异",
		House:      "已购房",
		Hokou:      "内蒙",
		Education:  "大学本科",
		Car:        "已购车",
	}

	save(expected)
}
