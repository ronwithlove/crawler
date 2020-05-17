package main

import (
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/persist"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

//itemserver的服务器
func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d",config.ItemSaverPort),config.ElasticIndex))
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