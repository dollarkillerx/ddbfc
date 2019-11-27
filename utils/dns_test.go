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
		err := DnsParsing(domain, time.Millisecond*400, 3)
		if err != nil {
			log.Println("err:  ",domain)
			continue
		}
	}
}
