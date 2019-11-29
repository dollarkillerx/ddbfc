/**
 * @Author: DollarKillerX
 * @Description: dns_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:00 2019/11/26
 */
package utils

import (
	"fmt"
	"github.com/bogdanovich/dns_resolver"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/miekg/dns"
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
		err := DnsParsing(domain, 1, 3)
		if err != nil {
			log.Println("err:  ", domain)
			log.Println(err)
			continue
		}
	}
}

// is new test
func TestNewDnsTest(t *testing.T) {
	domains := []string{
		"www.dollarkiller.com",
		"dollarkiller.com",
		"ps.cs",
		"xxxp.baidu.com",
		"www.baidu.com",
	}

	for _, domain := range domains {
		//testDomain(domain)
		parsing2 := DnsParsing2(domain, 1, 2)
		if parsing2 == nil {
			log.Println(domain)
		}
	}
}

func testDomain(domain string) {
	defer func() {
		log.Println("===============")
		fmt.Println()
	}()
	conn, e := newDns("8.8.8.8:53")
	if e != nil {
		panic(e)
	}

	// 进行dns查询
	msg := &dns.Msg{}
	// 进行A记录的查询
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	if err := conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
		panic(err)
	}

	if err := conn.WriteMsg(msg); err != nil {
		panic(err)
	}
	log.Println(domain)
	var err error

	if err := conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		panic(err)
	}
	if msg, err = conn.ReadMsg(); err != nil || len(msg.Question) == 0 {
		clog.PrintWa(err)
		clog.Println(domain)
		return
	}
	clog.Println(msg)

	record := NewRecord(domain, msg.Answer)
	if record == nil {
		clog.PrintEr("eee")
		return
	}
	clog.Println(record)
}

func TestTimeOUte(t *testing.T) {
	now := time.Now()
	add := now.Add(time.Second)
	//time.Sleep(time.Second)
	if time.Now().After(add) || time.Now().Equal(add) {
		log.Println("time out")
	}
}

func TestNewDns(t *testing.T) {
	resolver := dns_resolver.New([]string{"188.191.160.1"})
	ips, e := resolver.LookupHost("www.worldlink.com")
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(ips)
}
