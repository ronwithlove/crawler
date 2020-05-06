package persist

import (
	"context"
	"encoding/json"
	"github.com/crawler/crawler/model"
	"gopkg.in/olivere/elastic.v7"
	"testing"
)


func Test_save(t *testing.T) {
	expectd:=model.Profile{
		"3376375",
		"七……",
		"23",
		"未婚",
	}
	id,err:=save(expectd)//测试save方法

	if err!=nil{
		panic(err)
	}

	//这里还依赖于elastic 9200,如果全自动测试，最好用docker go client自己启动起来
	client, err := elastic.NewClient(
		elastic.SetSniff(false))//因为在docker里，没法sniff，所以关了
	if err!=nil{
		panic(err)
	}

	//拿出来
	resp,err:=client.Get().Index("dating_profile").Id(id).Do(context.Background())

	if err!=nil{
		panic(err)
	}

	t.Logf("%s",resp.Source)//这个t是testing的

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err!=nil{
		panic(err)
	}

	if actual!=expectd{
		t.Errorf("got %v;expected %v",actual,expectd)
	}
}