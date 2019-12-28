/**
 * @Author: DollarKillerX
 * @Description: bing_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午12:13 2019/12/7
 */
package datasource

import (
	"log"
	"testing"
)

func TestBing_ParseDomain(t *testing.T) {
	i := bingNew()
	domains, e := i.ParseDomain("google.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
