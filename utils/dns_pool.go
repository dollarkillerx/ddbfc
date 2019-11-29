/**
 * @Author: DollarKillerX
 * @Description: dns查询连接池
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:57 2019/11/27
 */
package utils

import (
	"errors"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/publicDns/service"
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

func GetDns() (*dns_resolver.DnsResolver, string) {
	dns := randomDns()
	resolver := dns_resolver.New([]string{dns})
	return resolver, dns
}

// dns相关的初始化 (获取全球dns)
func init() {
	// 获取全球public dns list
	lists, e := service.GetPublicDnsListService()
	if e != nil {
		clog.PrintWa(e)
		log.Fatalln("获取全球开公共dns失败")
	}

	// 更新dns列表
	for _, ic := range lists {
		dnsList = append(dnsList, ic.Ip)
	}
	log.Println("全球DnsList初始化成功")
}

// 负载均衡Dns
//func LoadDns() string {
//
//}
