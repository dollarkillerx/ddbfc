/**
 * @Author: DollarKillerX
 * @Description: ask_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:42 2019/12/6
 */
package datasource

import (
	"log"
	"testing"
)

func TestAsk_ParseDomain(t *testing.T) {
	i := askNew()
	domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(len(domains))
}

func TestAsk_ParseDomainParse(t *testing.T) {
	i := askNew()
	dom := "baidu.com"
	url := i.urlByPageNum(dom, 0)
	//log.Println(url)
	s, e := get(url)
	if e != nil {
		panic(e)
	}
	i.decode(dom, s)

}
