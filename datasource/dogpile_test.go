/**
 * @Author: DollarKillerX
 * @Description: dogpile_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:52 2019/12/7
 */
package datasource

import (
	"log"
	"testing"
)

func TestDogpile_ParseDomainUrl(t *testing.T) {
	i := dogpileNew()
	num := i.urlByPageNum("baidu.com", 1)
	log.Println(num)
}
