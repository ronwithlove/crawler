package persist

import (
	"context"
	"github.com/crawler/crawler/engine"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

func ItemSaver() chan engine.Item{
	out:=make(chan engine.Item)
	go func(){
		itemCount:=0
		for{
			item:=<-out
			log.Printf("Item saver: #%d: %v",itemCount,item)
			itemCount++

			 err := save(item)
			if err!=nil{
				log.Printf("Item Saver:error saving item %v: %v",item,err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))//因为在docker里，没法sniff，所以关了
	if err!=nil{
		return  err
	}

	indexService:=client.Index().Index("dating_profile").BodyJson(item)
	if item.Id!=""{//id是可续，如果可以从网页拿到，就加进去，拿不到，elastic会默认给
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())
	if err!=nil{
		return err
	}
	return nil
}

