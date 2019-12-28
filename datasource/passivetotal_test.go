/**
 * @Author: DollarKillerX
 * @Description: passivetotal_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:19 2019/12/7
 */
package datasource

import (
	"ddbf/utils/httplib"
	"log"
	"testing"
)

func TestPassivetotalUrl(t *testing.T) {
	lib, e := httplib.Get("https://api.passivetotal.org/v2/enrichment/subdomains")
	if e != nil {
		panic(e)
	}
	s, e := lib.Params("query", "baidu.com").
		Header("Content-Type", "application/json").
		ByteBigString()
	if e != nil {
		panic(e)
	}
	log.Println(s)
}
