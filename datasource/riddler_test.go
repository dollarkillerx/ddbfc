/**
 * @Author: DollarKillerX
 * @Description: riddler_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:59 2019/12/10
 */
package datasource

import (
	"log"
	"testing"
)

func TestRiddler_ParseDomain(t *testing.T) {
	r := riddlerNew()
	domains, e := r.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
