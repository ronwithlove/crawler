package persist

import (
	"context"
	"encoding/json"
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"gopkg.in/olivere/elastic.v7"
	"testing"
)


func Test_save(t *testing.T) {
	expectd:=engine.Item{
		Url:     "http://www.7799520.com/user/3376375.html",
		Id:      "3376375",
		Payload: model.Profile{
		"3376375",
		"七……",
		"23",
		"未婚",
		},
	}



	//这里还依赖于elastic 9200,如果全自动测试，最好用docker go client自己启动起来
	client, err := elastic.NewClient(
		elastic.SetSniff(false))//因为在docker里，没法sniff，所以关了
	if err!=nil{
		panic(err)
	}

	//1.先保存测试数据
	err=save(expectd)//测试save方法
	if err!=nil{
		panic(err)
	}

	//2.拿出来
	resp,err:=client.Get().Index("dating_profile").Id(expectd.Id).Do(context.Background())

	if err!=nil{
		panic(err)
	}

	t.Logf("%s",resp.Source)//这个t是testing的

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	//这两行把actual.Payload 改成mode.Profile格式
	actualProfile,_:=model.FromJsonObj(actual.Payload)
	actual.Payload=actualProfile

	//3.验证结果
	if actual!=expectd{
		t.Errorf("got %v;expected %v",actual,expectd)
	}
}