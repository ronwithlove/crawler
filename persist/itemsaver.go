package persist

import (
	"context"
	"github.com/crawler/crawler/engine"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

//index 是数据库名
func ItemSaver(index string) (chan engine.Item,error){
	client, err := elastic.NewClient(
		elastic.SetSniff(false))//因为在docker里，没法sniff，所以关了

	if err!=nil{
		return  nil,err
	}
	out:=make(chan engine.Item)
	go func(){
		itemCount:=0
		for{
			item:=<-out
			log.Printf("Item saver: #%d: %v",itemCount,item)
			itemCount++

			 err := Save(client,index,item)
			if err!=nil{
				log.Printf("Item Saver:error saving item %v: %v",item,err)
			}
		}
	}()
	return out,nil
}

//三个参数分别是：elastic客户端，表明，要存储的内容
func Save(client *elastic.Client, index string,item engine.Item) error {

	indexService:=client.Index().Index(index).BodyJson(item)
	if item.Id!=""{//id是可续，如果可以从网页拿到，就加进去，拿不到，elastic会默认给
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err!=nil{
		return err
	}
	return nil
}

