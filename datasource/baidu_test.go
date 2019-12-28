/**
 * @Author: DollarKillerX
 * @Description: baidu_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:54 2019/12/6
 */
package datasource

import (
	"log"
	"testing"
)

func TestBaidu_ParseDomainApiParse(t *testing.T) {
	//i := baiduNew()
	//i.decodeApi("baidu.com")
}

func TestBaidu_ParseDomainParse(t *testing.T) {
	i := baiduNew()
	dom := "google.com"
	url := i.urlByPageNum(dom, 0)
	log.Println(url)
	s, e := get(url)
	if e != nil {
		panic(e)
	}
	i.decode(dom, s)
}

func TestBaidu_ParseDomain(t *testing.T) {
	i := baiduNew()
	domains, e := i.ParseDomain("google.com")
	if e != nil {
		panic(e)
	}
	log.Println(len(domains))
}
