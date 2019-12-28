/**
 * @Author: DollarKillerX
 * @Description: google_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:35 2019/12/7
 */
package datasource

import (
	"log"
	"testing"
)

func TestGoogle_ParseDomainPpx(t *testing.T) {
	g := googleNew()
	num := g.urlByPageNum("dollarkiller.com", 0, 1)
	log.Println(num)
}

func TestGoogle_ParseDomainDecode(t *testing.T) {
	g := googleNew()
	url := g.urlByPageNum("dollarkiller.com", 0, 0)
	s, e := get(url)
	if e != nil {
		panic(e)
	}
	log.Println(url)
	g.decode("dollarkiller.com", s)
}

func TestGoogle_ParseDomain(t *testing.T) {
	g := googleNew()
	domains, e := g.ParseDomain("dollarkiller.com")
	//domains, e := g.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
