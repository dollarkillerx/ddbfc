/**
 * @Author: DollarKillerX
 * @Description: goo_ne_jp_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:26 2019/12/9
 */
package datasource

import (
	"log"
	"testing"
)

func TestGooNeJp_ParseDomain(t *testing.T) {
	jpNew := gooNeJpNew()
	domains, e := jpNew.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
