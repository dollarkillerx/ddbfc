/**
 * @Author: DollarKillerX
 * @Description: dns_pool_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:35 2019/11/30
 */
package utils

import (
	"github.com/bogdanovich/dns_resolver"
	"log"
	"testing"
)

func TestDnsNum(t *testing.T) {
	t.Log(len(DnsSource))
	resolver := dns_resolver.New([]string{"8.8.8.8"})
	ips, e := resolver.LookupHost("www.dollarkiller.com")
	if e != nil {
		panic(e)
	}
	log.Println(ips)

}

func TestDnsPol(t *testing.T) {
	get := SDnsGet()
	get.DnsIds.ErrNum++
	get.DnsIds.ErrNum++

	for k, v := range DnsSource {
		if k == get.DnsIds.ID {
			log.Println(v)
		}
	}
}
