/**
 * @Author: DollarKillerX
 * @Description: crtsh_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:47 2019/12/6
 */
package datasource

import (
	"log"
	"testing"
)

func TestCrtsh_ParseDomain(t *testing.T) {
	i := crtshNew()
	//domains, e := i.ParseDomain("baidu.com")
	domains, e := i.ParseDomain("dollarkiller.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
