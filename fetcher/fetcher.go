package fetcher

import (
	"bufio"
	"fmt"
	"github.com/crawler/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter=time.Tick(time.Second/config.Qps)
//传入url，传出文本
func Fetch(url string)([]byte,error){
	<-rateLimiter//通过这个channel来降低速度
	log.Printf("Fetching url %s",url)
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
	//bodyReader:=bufio.NewReader(resp.Body)
	//e:=determineEncoding(bodyReader)
	//utf8Reader:=transform.NewReader(bodyReader,e.NewDecoder())
	//return ioutil.ReadAll(utf8Reader)

	return ioutil.ReadAll(resp.Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err:=r.Peek(1024)
	if err!=nil{
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}