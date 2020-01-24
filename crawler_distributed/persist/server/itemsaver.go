package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"learnCrawler/crawler_distributed/config"
	"learnCrawler/crawler_distributed/persist"
	"learnCrawler/crawler_distributed/rpcsupport"
	"log"
)

var port = flag.Int("port",0,"the port for me to listen on")
func main() {
	flag.Parse()
	if *port==0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRpc(fmt.Sprintf(":%d",*port), config.ElasticIndex))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
