package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

func ItemSaver() chan interface{}{
	out:=make(chan interface{})
	go func(){
		itemCount:=0
		for{
			item:=<-out
			log.Printf("Item saver: #%d: %v",itemCount,item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{})(id string,err error){
	client, err := elastic.NewClient(
		elastic.SetSniff(false))//因为在docker里，没法sniff，所以关了
	if err!=nil{
		return "",err
	}

	resp, err := client.Index().Index("dating_profile").BodyJson(item).Do(context.Background())
	if err!=nil{
		return "",err
	}
	return  resp.Id, nil
}

