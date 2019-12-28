/**
 * @Author: DollarKillerX
 * @Description: httplib_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午7:26 2019/12/4
 */
package httplib

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestStringSP(t *testing.T) {
	tag := "https://www.ok.com?asd=sad"
	index := strings.Index(tag, "?")
	log.Println(tag[:index+1])
}

func TestGet(t *testing.T) {
	lib, e := Get("http://www.baidu.com")
	if e != nil {
		panic(e)
	}
	s, e := lib.String()
	if e != nil {
		panic(e)
	}
	log.Println(s)
}

func TestPost(t *testing.T) {

}

func TestBigRead(t *testing.T) {
	file, e := os.Open("user_agent.go")
	if e != nil {
		panic(e)
	}
	bytes, e := readBigFile(file)
	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("cc.txt", bytes, 00666)
}

func readBigFile(os io.ReadCloser) ([]byte, error) {
	// 创建接受容器
	db := make([]byte, 0)

	// 创建读取器
	reader := bufio.NewReader(os)
	// 创建读取缓冲区
	buf := make([]byte, 1024)
	for {
		n, e := reader.Read(buf)
		if e != nil {
			if e == io.EOF {
				break
			}
			return nil, e
		}
		db = append(db, buf[:n]...)
	}
	return db, nil
}

func TestHttp(t *testing.T) {
	url := "http://speedtest-sfo1.digitalocean.com/5gb.test"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	log.Println(resp.ContentLength)
}

//表示头500个字节：bytes=0-499
//表示第二个500字节：bytes=500-999
//表示最后500个字节：bytes=-500
//表示500字节以后的范围：bytes=500-
//第一个和最后一个字节：bytes=0-0,-1

// 分段下载测试
func TestHcc(t *testing.T) {
	//url := "http://speedtest-sfo1.digitalocean.com/5gb.test"
	url := "https://index.commoncrawl.org/CC-MAIN-2008-2009-index?filter=%3Dstatus%3A200&fl=url%2Cstatus&output=json&pageSize=2000&url=%2A.baidu.com"

	httpClient := &http.Client{}
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		panic(e)
	}
	req.Header.Set("User-Agent", getRandUA())
	req.Header.Set("Range", "bytes=1-500")
	response, e := httpClient.Do(req)
	if e != nil {
		panic(e)
	}
	defer response.Body.Close()

	if response.StatusCode == 206 {
		log.Println("支持断点下载")
	} else {
		log.Println(response.StatusCode)
		log.Println("不支持断点下载")
	}
}
