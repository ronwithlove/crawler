package parser

import (
	"github.com/crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err!=nil{
		panic(err)
	}

	result:=ParseProfile(contents,"3376375")//这个userid是网页里的

	if len(result.Items)!=1{
		t.Errorf("Items应该只有1个元素，现在有 %v",len(result.Items))
	}

	profile:=result.Items[0].(model.Profile)

	expectd:=model.Profile{
		"3376375",
		"七……",
		"23",
		"未婚",
	}

	if profile!=expectd{
		t.Errorf("expected %v:,but was %v.",expectd,profile)
	}
}
