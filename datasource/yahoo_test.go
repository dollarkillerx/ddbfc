/**
 * @Author: DollarKillerX
 * @Description: yahoo_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:21 2019/12/5
 */
package datasource

import (
	"log"
	"testing"
)

func TestYahoo_ParseDomain(t *testing.T) {
	y := yahooNew()
	//y.ParseDomain("bikccsad.com")
	domains, e := y.ParseDomain("google.com")
	//domains, e := y.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
