package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err!=nil{
		panic(err)
	}

	//fmt.Printf("%s\n",contents)
	result:=ParseCity(contents,"")

	const resultSize=15
	expectedUrls:=[]string{
		"http://www.7799520.com/user/3376112.html",
		"http://www.7799520.com/user/3375292.html",
		"http://www.7799520.com/user/3372945.html",
	}

	if len(result.Requests)!=resultSize{
		t.Errorf("expected %d requests; but had %d",resultSize, len(result.Requests))
	}
	//校对前3个url是否正确
	for i,url:=range expectedUrls{
		if result.Requests[i].Url!=url{
			t.Errorf("expected url #%d: %s; but was %s", i, url,result.Requests[i].Url)
		}
	}

}