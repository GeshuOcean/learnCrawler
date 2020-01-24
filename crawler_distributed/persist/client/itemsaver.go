package client

import (
	"learnCrawler/crawler/engine"
	"learnCrawler/crawler_distributed/config"
	"learnCrawler/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item,  error){
	client,err := rpcsupport.NewClient(host)
	if err!=nil{
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item "+"#%d: %v", itemCount, item)
			itemCount++

			//Call RPC to save item
			result :=""
			err:=client.Call(config.ItemSaverRpc,item,&result)
			if err != nil {
				log.Print("Item Server:item:%v, err:%v", item, err)
			}
		}
	}()
	return out,nil
}
