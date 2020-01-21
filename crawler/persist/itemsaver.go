package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"learnCrawler/crawler/engine"
	"log"
)

func ItemServer() (chan engine.Item,error) {
	client, err := elastic.NewClient(
		//must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return  nil,err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item #%d:%v", itemCount, item)
			itemCount++

			err := save(client,item)
			if err != nil {
				log.Printf("Item Saver:error saving item %v:%v", item, err)
			}

		}
	}()
	return out,nil
}

func save(client *elastic.Client,item engine.Item) (err error) {
	if item.Type==""{
		return errors.New("must supply Type")
	}
	indexService :=client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id!=""{
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return  err
	}
	return nil
}
