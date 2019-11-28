/**
 * @Author: DollarKillerX
 * @Description: dns解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:50 2019/11/26
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/miekg/dns"
	"net"
	"time"
)

// dns解析测试
// @param 域名
// @param 超时时间
// @param 尝试次数
func DnsParsing(domain string, timeout int, tryNum int) error {
	var err error
	var ips []net.IP
	for i := 0; i < tryNum; i++ {
		dns := GetDns()
		ips, err = dns.LookupHost(domain)
		if err == nil && len(ips) != 0 {
			return nil
		} else if err != nil {
			if checkTimeOut(err) {
				return TimeOut
			}
		}
	}
	if err == nil {
		return errors.New("not dns")
	}
	return err
}

var TimeOut = errors.New("timeout")

func DnsParsing2(domain string, timeout int, tryNum int) error {
	var err error
	for i := 0; i < tryNum; i++ {
		timeo := time.Now().Add(time.Second)
		conn, er := newDns("8.8.8.8:53")
		if er != nil {
			continue
		}
		// 进行dns查询
		msg := &dns.Msg{}
		// 进行A记录的查询
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
		if err = conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
			clog.PrintEr(err)
			continue
		}

		if err = conn.WriteMsg(msg); err != nil {
			return TimeOut
			continue
		}

		if err = conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
			clog.PrintEr(err)
			continue
		}
		if msg, err = conn.ReadMsg(); err != nil || len(msg.Question) == 0 {
			if err != nil {
				if checkTimeOut(err) {
					return TimeOut
				}
			}
			continue
		}

		if time.Now().After(timeo) || time.Now().Equal(timeo) {
			return TimeOut
		}

		record := NewRecord(domain, msg.Answer)
		if record != nil {
			return nil
		}

	}
	if err == nil {
		err = errors.New("dns err")
	}
	return err
}

// 创建dns连接
func newDns(host string) (*dns.Conn, error) {
	return dns.DialTimeout("udp", host, time.Second)
}

// 旧的版本
//func DnsParsing(domain string, timeout time.Duration, tryNum int) error {
//	var err error
//	for i := 0; i < tryNum; i++ {
//		wt, _ := context.WithTimeout(context.TODO(), timeout)
//		host, err := net.DefaultResolver.LookupHost(wt, domain)
//		//host, err := net.LookupHost(domain)
//		if err == nil && host != nil {
//			return nil
//		}
//	}
//	if err == nil {
//		return errors.New("not dns")
//	}
//	return err
//}

//func Dnsc() {
//	client := dns.Client{
//		Timeout: 5 * time.Second,
//	}
//
//
//}

// 更新的Dns

// Record DNS 记录
type Record struct {
	Domain string
	Type   string
	Target string
	IP     []string
}

// NewRecord 新建 DNS 记录
func NewRecord(domain string, response []dns.RR) *Record {
	if len(response) == 0 || isPanDNS(domain, response) {
		return nil
	}

	record := Record{Domain: domain}
	switch firstAnswer := response[0].(type) {
	case *dns.CNAME:
		record.Type = "CNAME"
		record.Target = trimSuffixPoint(firstAnswer.Target)
		response = response[1:]
	case *dns.A:
		record.Type = "A"
	default:
		return nil
	}

	for _, ans := range response {
		if a, ok := ans.(*dns.A); ok {
			record.IP = append(record.IP, a.A.String())
		}
	}

	return &record
}

var panDNSBlackList = map[string][]string{}

var dnsServerAddress string

func init() {
	dnsServerAddress = "8.8.8.8:53"
}

// queryPanDNS 生成父级域名泛解析黑名单
func queryPanDNS(domain string) (firstTime bool) {
	// 如果父级域名已存在，不再查询
	if _, ok := panDNSBlackList[domain]; ok {
		return
	}

	// md5 域名
	hash := md5.New()
	hash.Write([]byte(domain))
	md5Domain := hex.EncodeToString(hash.Sum(nil))[8:24] + "." + domain

	msg := &dns.Msg{}
	msg.SetQuestion(dns.Fqdn(md5Domain), dns.TypeA)
	in, err := dns.Exchange(msg, dnsServerAddress)
	if err != nil || len(in.Answer) == 0 {
		return
	}

	var rr string
	for _, ans := range in.Answer {
		switch ans := ans.(type) {
		case *dns.CNAME:
			rr = ans.Target
		case *dns.A:
			rr = ans.A.String()
		}
		panDNSBlackList[domain] = append(panDNSBlackList[domain], rr)
	}

	return true
}

// 判断是否是泛解析
func isPanDNS(domain string, response []dns.RR) bool {
	pd := parentDomain(domain)
	firstTime := queryPanDNS(pd)

	// 第一次探测该父级域名，不判定是否是泛解析
	if firstTime {
		return false
	}

	// 无记录，不是泛解析
	records, ok := panDNSBlackList[pd]
	if !ok {
		return false
	}

	// 存在记录，且 CNAME/IP 均存在于黑名单中，是泛解析
	var rr string
	for _, ans := range response {
		switch ans := ans.(type) {
		case *dns.CNAME:
			rr = ans.Target
		case *dns.A:
			rr = ans.A.String()
		}
		if !strInSlice(rr, records) {
			return false
		}
	}

	return true
}

func checkTimeOut(err error) bool {
	//if index := strings.Index(err.Error(), "timeout"); index != -1 {
	//	return true
	//}
	//return false
	if err.Error() != "NXDOMAIN" {
		return true
	}
	return false
}
