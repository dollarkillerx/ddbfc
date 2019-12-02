/**
 * @Author: DollarKillerX
 * @Description: dns_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:00 2019/11/26
 */
package utils

import (
	"log"
	"testing"
	"time"
)

func TestDnsParsing(t *testing.T) {
	domains := []string{
		"www.baidu.com",
		"www.dollarkiller.com",
		"dollarkiller.com",
		"qq.com",
		"google.com",
		"1688.com",
		"360.com",
	}

	for _, domain := range domains {
		coon, e := GetDnsByPool(time.Second)
		if e != nil {
			panic(e)
		}
		b, e := coon.DnsParse(domain)
		ec := ReleaseDns(coon)
		if ec != nil {
			log.Fatalln(ec)
		}
		if e != nil {
			log.Fatalln(e)
		}
		if b {
			log.Println(domain)
		}
	}
}
