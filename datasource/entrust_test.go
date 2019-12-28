/**
 * @Author: DollarKillerX
 * @Description: entrust_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午6:09 2019/12/5
 */
package datasource

import (
	"ddbf/utils/httplib"
	"log"
	"testing"
)

func TestEntrust_ParseDomain2(t *testing.T) {
	i := entrustNew()
	domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}

func TestEntrust_ParseDomain(t *testing.T) {
	dom := "baidu.com"
	i := entrustNew()
	url := i.getURL(dom)
	lib, _ := httplib.Get(url)
	s, _ := lib.String()

	domains, e := i.decode(dom, s)
	if e != nil {
		panic(e)
	}
	log.Println(domains)
	log.Println(len(domains))
}

func TestEntrust_ParseDomainUrl(t *testing.T) {
	i := entrustNew()
	url := i.getURL("dollarkiller.com")
	log.Println(url)
}
