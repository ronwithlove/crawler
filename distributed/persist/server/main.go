package main

import (
	"github.com/crawler/crawler/distributed/persist"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234","dating_profile"))
}

func serveRpc(host,index string)error{
	client,err:=elastic.NewClient(elastic.SetSniff(false))
	if err!=nil{
		return err
	}
	return rpcsupport.ServeRpc(host,&persist.ItemSaverService{
		Client: client,
		Index:  index,
	})//实例化serve

}