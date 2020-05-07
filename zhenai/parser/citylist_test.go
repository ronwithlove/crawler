package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//如果网页内容不能访问，那么test也没法继续，所以事先把网页拷贝下来，直接去读文件来解决这个问题
	//测试一般都用这种做法
	//contents, err := fetcher.Fetch("http://city.7799520.com")
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err!=nil{
		panic(err)
	}

	//fmt.Printf("%s\n",contents)
	result:=ParseCityList(contents)

	const resultSize=389 //真爱是470
	expectedUrls:=[]string{
		"http://city.7799520.com/anhui",
		"http://city.7799520.com/aomen",
		"http://city.7799520.com/anqing",
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