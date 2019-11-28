/**
 * @Author: DollarKillerX
 * @Description: dns查询连接池
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:57 2019/11/27
 */
package utils

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/bogdanovich/dns_resolver"
)

var dnsList = []string{
	"1.1.1.1",
	"1.0.0.1",
	"8.8.8.8",
	"8.8.4.4",
	"9.9.9.9",
	"9.9.9.10",
	//"208.67.222.222",
	//"208.67.220.220",
	//"8.26.56.26",
	//"8.20.247.20",
	//"64.6.64.6",
	//"64.6.65.6",
}

// dns池
type dnsPool struct {
	bufChan chan *dns_resolver.DnsResolver
}

var DnsPool *dnsPool

// 初始化dns查询连接池
func init() {
	// 初始化连接池 容量100  单个查询超时5
	DnsPool = NewDnsPool(1, 1)
}

func NewDnsPool(num int, timeout int) *dnsPool {
	pool := dnsPool{}
	pool.bufChan = make(chan *dns_resolver.DnsResolver, num)

	// 创建对象
	for i := 0; i < num; i++ {
		resolver := dns_resolver.New([]string{randomDns()})
		resolver.RetryTimes = timeout
		pool.bufChan <- resolver
	}

	return &pool
}

// 获取dns连接对象 (这里没有作超时控制,就让他阻塞)
func (d *dnsPool) GetDns() *dns_resolver.DnsResolver {
	select {
	case dns := <-d.bufChan:
		return dns
	}
}

// 放回对象
func (d *dnsPool) ReleaseDns(dns *dns_resolver.DnsResolver) error {
	select {
	case d.bufChan <- dns:
		log.Println("被放回了")
		log.Println(len(d.bufChan))
		return nil
	default:
		return errors.New("pool overflow")
	}
}

// 获取随即dns
func randomDns() string {
	i := len(dnsList)
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(i)
	return dnsList[intn]
}

func GetDns() *dns_resolver.DnsResolver {
	resolver := dns_resolver.New([]string{randomDns()})
	return resolver
}
