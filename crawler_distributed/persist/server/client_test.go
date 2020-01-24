package main

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/model"
	"learnCrawler/crawler_distributed/config"
	"learnCrawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	const host = ":1234"
	//start ItemSaverServer
	go serverRpc(host, "test1")
	time.Sleep(time.Second)

	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "www.sa.com",
		Type: "zhaiai",
		Id:   "111111",
		Payload: model.Profile{
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
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s;err:%s", result, err)
	}
}
