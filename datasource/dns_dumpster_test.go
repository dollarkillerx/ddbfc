/**
 * @Author: DollarKillerX
 * @Description: dns_dumpster_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:12 2019/12/6
 */
package datasource

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestDnsDumpster_ParseDomain(t *testing.T) {
	dump := dnsDumpsterNew()
	domains, e := dump.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}

func TestDnsDumpster_ParseDomainPost(t *testing.T) {
	dump := dnsDumpsterNew()
	err := dump.getCsrfToken()
	if err != nil {
		panic(err)
	}

	s, err := dump.postFrom("baidu.com")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("dumoster.html", []byte(s), 00666)
	if err != nil {
		panic(err)
	}
}
