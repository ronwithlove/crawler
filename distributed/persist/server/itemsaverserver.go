package main

import (
	"flag"
	"fmt"
	"github.com/crawler/crawler/config"
	"github.com/crawler/crawler/distributed/persist"
	"github.com/crawler/crawler/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

//itemserver的服务器
var port=flag.Int("port",0,"open port to listen")

func main()  {
	flag.Parse()
	if *port==0{
		fmt.Println("must specify a port")
		return
	}
	//使用命令行键入port
	log.Fatal(serveRpc(fmt.Sprintf(":%d",*port),config.ElasticIndex))
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