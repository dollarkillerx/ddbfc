/**
 * @Author: DollarKillerX
 * @Description: exalead_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:27 2019/12/7
 */
package datasource

import (
	"log"
	"testing"
)

func TestExalead_ParseDomain(t *testing.T) {
	i := exaleadNew()
	domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
