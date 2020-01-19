package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
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

	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	//TODO Try to start up elastic search
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	json.Unmarshal(*resp.Source, &actual)
	if actual!=expected{

	}
}
