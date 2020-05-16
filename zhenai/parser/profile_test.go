package parser

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err!=nil{
		panic(err)
	}

	result:=parseProfile(contents,"http://www.7799520.com/user/3376375.html","3376375")//这个userid是网页里的

	if len(result.Items)!=1{
		t.Errorf("Items应该只有1个元素，现在有 %v",len(result.Items))
	}

	actual:=result.Items[0]

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

	if actual!=expectd{
		t.Errorf("expected %v:,but was %v.",expectd,actual)
	}
}
