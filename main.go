package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
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

	//网页转码方法一
	//这个包在golang.org/x/text里  encoding 下
	//如果网页是GBK，需要转码，要不然打出来是乱码
	//utf8Reader:=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	//all,err:=ioutil.ReadAll(utf8Reader)

	//网页转码方法二
	//这个需要用到golang.org/x/net下的html
	//e:=determineEncoding(resp.Body)
	//utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	//all,err:=ioutil.ReadAll(utf8Reader)

	all,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",all)

}


func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err!=nil{
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}