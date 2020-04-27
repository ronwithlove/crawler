package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//传入url，传出文本
func Fetch(url string)([]byte,error){
	resp,err:=http.Get(url)
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		return nil,
		fmt.Errorf("worng status code: %d",resp.StatusCode)
	}

	//如果网页编码不是utf8就需要先转码
	//e:=determineEncoding(resp.Body)
	//utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	//return ioutil.ReadAll(utf8Reader)

	return ioutil.ReadAll(resp.Body)
}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err!=nil{
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}