/**
 * @Author: DollarKillerX
 * @Description: dns解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:50 2019/11/26
 */
package utils

import (
	"errors"
	"github.com/miekg/dns"
	"net"
	"time"
)

// dns解析测试
// @param 域名
// @param 超时时间
// @param 尝试次数
func DnsParsing(domain string, timeout time.Duration, tryNum int) error {
	var err error
	for i := 0; i < tryNum; i++ {
		//wt, _ := context.WithTimeout(context.TODO(), timeout)
		//_, err = net.DefaultResolver.LookupHost(wt, domain)
		host, err := net.LookupHost(domain)
		if err == nil && host != nil {
			return nil
		}
	}
	if err == nil {
		return errors.New("not dns")
	}
	return err
}

//func Dnsc() {
//	client := dns.Client{
//		Timeout: 5 * time.Second,
//	}
//
//
//}

func CNAME2(src string, dnsService string) (dst []string, err error) {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	var lastErr error
	// retry 3 times
	for i := 0; i < 3; i++ {
		m := dns.Msg{}
		// 最终都会指向一个ip 也就是typeA, 这样就可以返回所有层的cname.
		m.SetQuestion(src+".", dns.TypeA)
		r, _, err := c.Exchange(&m, dnsService+":53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}

		dst = []string{}
		for _, ans := range r.Answer {
			record, isType := ans.(*dns.CNAME)
			if isType {
				dst = append(dst, record.Target)
			}
		}
		lastErr = nil
		break
	}

	err = lastErr

	return
}
