package main

import (
	"fmt"
	"learnCrawler/crawler_distributed/config"
	"learnCrawler/crawler_distributed/rpcsupport"
	"learnCrawler/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T){
	//启动服务端
	const host=":9000"
	go rpcsupport.ServerRpc(host,worker.CrawlService{})
	time.Sleep(time.Second)

	//客户端传参调用
	client,err:=rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}
	req:=worker.Request{
		Url:    "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"安静的雪",
		},
	}
	var result worker.ParseResult
	err=client.Call(config.CrawlServiceRpc,req,&result)
	if err!=nil{
		t.Error(err)
	}else{
		fmt.Println(result)
	}
}
