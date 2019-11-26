/**
 * @Author: DollarKillerX
 * @Description: ping_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:26 2019/11/26
 */
package test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	domains := []string{
		"baidu.com",
		"dollarkiller.com",
		"ppc.c",
	}

	for _, v := range domains {
		timeout, _ := context.WithTimeout(context.TODO(), time.Millisecond*200)
		ns, err := net.DefaultResolver.LookupHost(timeout, v)
		if err != nil {
			log.Println("err == = = = = == ")
			log.Println(err)
			continue
		}

		log.Println(ns)
	}
}
