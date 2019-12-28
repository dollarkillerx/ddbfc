/**
 * @Author: DollarKillerX
 * @Description: threatminer_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:15 2019/12/9
 */
package datasource

import (
	"log"
	"testing"
)

func TestThreatminer_ParseDomain(t *testing.T) {
	tc := threatminerNew()
	//domains, e := tc.ParseDomain("baidu.com")
	domains, e := tc.ParseDomain("dollarkiller.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
