/**
 * @Author: DollarKillerX
 * @Description: fireball_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:53 2019/12/9
 */
package datasource

import (
	"log"
	"testing"
)

func TestFireball_ParseDomain(t *testing.T) {
	i := fireballNew()
	domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
