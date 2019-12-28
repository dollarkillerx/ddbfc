/**
 * @Author: DollarKillerX
 * @Description: googlect_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:49 2019/12/7
 */
package datasource

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestGooglect_ParseDomainGetUrl(t *testing.T) {
	i := googlectNew()
	dnsurl := i.getDNSURL("baidu.com", "")
	log.Println(dnsurl)
}

func TestGooglect_ParseDomainDow(t *testing.T) {
	i := googlectNew()
	dnsurl := i.getDNSURL("dollarkiller.com", "")
	s, e := i.http(dnsurl)
	if e != nil {
		panic(e)
	}
	e = ioutil.WriteFile("googlect.json", []byte(s), 00666)
	if e != nil {
		panic(e)
	}
}
