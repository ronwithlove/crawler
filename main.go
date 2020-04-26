package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
	resp,err:=http.Get(
		"http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}

	all,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte){
	re:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _,m:=range matches{
		fmt.Printf("城市：%s，URL: %s\n",m[2],m[1])//用空格把每个元素分开
	}
	fmt.Printf("找到：%d 个结果\n",len(matches))
}