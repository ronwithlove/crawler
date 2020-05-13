package main

import (
	"github.com/crawler/crawler/distributed/rpcsupport"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host=":1234"
	//1.start ItemSaverServer
	go serveRpc(host,"test1")
	time.Sleep(time.Second)

	//2.start ItemSaverClient
	client,err:=rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}

	//3.Call save
	item:=engine.Item{
		Url:     "http://www.7799520.com/user/3376375.html",
		Id:      "3376375",
		Payload: model.Profile{
			"3376375",
			"七……",
			"23",
			"未婚",
		},
	}
	result:=""
	err=client.Call("ItemSaverService.Save",item, &result)
	if err!=nil||result!="ok"{
		t.Errorf("result:%s;err:%s",result,err)
	}
}
