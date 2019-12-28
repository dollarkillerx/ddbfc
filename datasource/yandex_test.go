/**
 * @Author: DollarKillerX
 * @Description: yandex_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:21 2019/12/9
 */
package datasource

import (
	"log"
	"testing"
)

func TestYandex_ParseDomainuRL(t *testing.T) {
	dom := "dollarkiller.com"
	i := yandexNew()
	url := i.getUrl(dom, 0)
	log.Println(url)
}

func TestYandex_ParseDomain(t *testing.T) {
	dom := "dollarkiller.com"
	i := yandexNew()
	domains, e := i.ParseDomain(dom)
	if e != nil {
		panic(e)
	}
	log.Println(domains)
}
