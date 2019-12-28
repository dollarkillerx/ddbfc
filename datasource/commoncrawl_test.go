/**
 * @Author: DollarKillerX
 * @Description: commoncrawl_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:09 2019/12/5
 */
package datasource

import (
	"ddbf/utils/httplib"
	"log"
	"strings"
	"testing"
)

func TestCommoncrawl_ParseDomainUrlParser(t *testing.T) {
	c := "baidu.com"
	url := "http://www.baidu.com/baidu?word=XML+Weather+Docklet&tn=max2_cb"
	index := strings.Index(url, "://")
	end := strings.Index(url, c)
	log.Println(url[index+3 : end+len(c)])
}

func TestCommoncrawl_ParseDomain(t *testing.T) {
	i := commoncrawlNew()
	domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}

func TestOp(t *testing.T) {
	url := "https://index.commoncrawl.org/CC-MAIN-2008-2009-index?filter=%3Dstatus%3A200&fl=url%2Cstatus&output=json&pageSize=2000&url=%2A.baidu.com"

	lib, e := httplib.Get(url)
	if e != nil {
		panic(e)
	}
	s, e := lib.ByteBigString()
	if e != nil {
		panic(e)
	}

	log.Println(s)
}
