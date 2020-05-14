package persist

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/persist"
	"gopkg.in/olivere/elastic.v7"
	"log"
)

type ItemSaverService struct{
	Client *elastic.Client
	Index string
}

//写一个服务器的方法，可以让服务器调用，其中就用了原来的Save方法
//这里reslut随便写一个，因为要满足rpc服务的格式:传一个，返一个
func(s *ItemSaverService)Save(item engine.Item,result *string)error{
	err := persist.Save(s.Client, s.Index, item)//调用Save方法
	log.Printf("Item %v saved.",item)
	if err==nil{
		*result="ok"
	}else{
		log.Printf("Error saving item %v: %v",item,err)
	}
	return err
}