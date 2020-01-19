package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code:%d", resp.StatusCode)
	}

	bodyRead:=bufio.NewReader(resp.Body)
	//识别响应的编码
	e := determinEncoding(bodyRead)
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(bodyRead, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}


//判断返回响应的编码，默认UTF8
func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
