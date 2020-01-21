package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler/model"
	"testing"
)

func TestItemServer(t *testing.T) {
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/105466768",
		Type: "zhenai",
		Id:   "105466768",
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
			Car:        "已购车",
		},
	}


	//TODO Try to start up elastic search
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = save(client,index,expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)
	actualProfile,_:=model.FromJsonObj(actual.Payload)
	actual.Payload=actualProfile
	if actual != expected {

	}
}
