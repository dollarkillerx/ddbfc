/**
 * @Author: DollarKillerX
 * @Description: dns查询连接池
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:57 2019/11/27
 */
package utils

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/bogdanovich/dns_resolver"
	"github.com/dollarkillerx/easyutils/httplib"
)

var dnsList = []string{
	//全球dns
	"1.1.1.1",
	"1.0.0.1",
	"8.8.8.8",
	"8.8.4.4",
	"9.9.9.9",
	"9.9.9.10",
	"208.67.222.222",
	"208.67.220.220",
	"4.2.2.1",
	"4.2.2.2",
	"77.88.8.8",
	"77.88.8.7",
	"77.88.8.1",
	"77.88.8.2",
	"77.88.8.3",
	"64.6.64.6",
	"64.6.65.6",
	"8.26.56.26",
	"8.20.247.20",
	"84.200.69.80",
	"84.200.70.40",

	// 区域dns 国内
	"1.2.4.8",
	"202.98.0.68",
	"119.29.29.29",
	"114.114.114.114",
	"114.114.114.115",
	"223.5.5.5",
	"223.6.6.6",
	"180.76.76.76",

	// 区域性的第三方dns解析
	"101.101.101.101",
	"101.102.103.104",
	"112.121.178.187",
	"168.95.192.1",
	"203.112.2.4",
	"168.126.63.1",
	"168.126.63.2",
	"80.80.80.80",
	"80.80.81.81",
	"216.146.35.35",
	"216.146.36.36",
	"74.82.42.42",
	"195.46.39.39",
	"195.46.39.40",
	"109.69.8.51",
	"156.154.70.1",
	"156.154.71.1",
	"37.235.1.174",
	"37.235.1.177",

	// 国外运营商的dns

	// at
	"194.150.168.168",
	// jd
	"210.228.48.29",
	// uk
	"195.99.66.220",
	// fr
	"194.98.65.65",
	// ca
	"142.103.1.1",
	// 瑞典
	"195.67.27.18",
	// 韩国
	"168.126.63.1",
	// 瑞士
	"195.186.1.111",
}
var DnsSourceMu sync.Mutex
var DnsSource = []*DnsServer{}

type DnsServer struct {
	IP     string // 服务器ip
	RunNum int    // 运行次数
	ErrNum int    // 失败次数
	ID     int    // 此dns的下标

	Tag int // 标识位 0在正常1在休息
	//sync.Mutex     // 锁
}

// 新的策略每100次解析休息100ms  每此失败休息200ms 如果失败数>10 休息3s  > 15 休息3min

// 初始化DnsSource
func init() {
	// 验证dns有效性
	for _, ip := range dnsList {
		resolver := dns_resolver.New([]string{ip})
		ips, e := resolver.LookupHost("www.bing.com")
		if e != nil || len(ips) == 0 {
			continue
		}
		dns := &DnsServer{
			IP:     ip,
			RunNum: 0,
			ErrNum: 0,
			ID:     0,
		}
		DnsSourceMu.Lock()
		DnsSource = append(DnsSource, dns)
		DnsSourceMu.Unlock()
	}
	log.Println("dns列表初始化完毕")
	log.Println("有效dns数量: ", len(DnsSource))
}

func sDnsGet() (*DnsServer, int) {
	DnsSourceMu.Lock()
	lang := len(DnsSource)
	DnsSourceMu.Unlock()
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(lang)

	DnsSourceMu.Lock()
	dns := DnsSource[intn]
	DnsSourceMu.Unlock()
	return dns, intn
}

// 获取 dnsSource中的dns
func SDnsGet() *CheckDns {
	for {
		dns, intn := sDnsGet()

		if dns.Tag == 0 {
			defer func() {
				dns.Tag = 0
			}()
			// 限流等操作
			if dns.RunNum > 100 {
				dns.Tag = 1
				time.Sleep(time.Millisecond * 200)
				dns.RunNum = 0
			}

			if dns.ErrNum > 10 && dns.ErrNum < 15 {
				dns.Tag = 1
				time.Sleep(time.Second * 3)
				dns.ErrNum = 0
			} else {
				dns.Tag = 1
				time.Sleep(time.Millisecond)
				dns.ErrNum = 0
			}

			dns.ID = intn

			return &CheckDns{
				Dns:    dns_resolver.New([]string{dns.IP}),
				DnsIds: dns,
			}
		}
		continue
		select {}
	}
}

// dns通用返回  dns，标签 0 正常 1 错误
func dnsSeed(dns *CheckDns, tag int) {
	switch tag {
	case 0:
		dns.DnsIds.RunNum++
	case 1:
		dns.DnsIds.RunNum++
		dns.DnsIds.ErrNum++
	}
}

type CheckDns struct {
	Dns    *dns_resolver.DnsResolver
	DnsIds *DnsServer
}

func (c *CheckDns) Check(host string) ([]net.IP, error) {
	ips, e := c.Dns.LookupHost(host)
	if e != nil {
		switch e.Error() {
		case "NXDOMAIN": // 域名没有
			dnsSeed(c, 0)
			return ips, e
		case "SERVFAIL": // 服务失败
			dnsSeed(c, 1)
		default: // 其他错误
			dnsSeed(c, 1)
		}
	}
	dnsSeed(c, 0)
	return ips, e
}

// 以下是旧的策略

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
//func init() {
//	// 获取全球public dns list
//	lists, e := getDnsList()
//	if e != nil {
//		clog.PrintWa("DnsList 获取失败")
//		log.Fatalln(e)
//	}
//	// 更新dns列表
//	for _, ic := range lists {
//		dnsList = append(dnsList, ic)
//	}
//	log.Println("全球DnsList初始化成功")
//}

func getDnsList() ([]string, error) {
	//https://dns.bilibilil.cf/getdnslist
	bytes, e := httplib.EuUserGet("https://dns.bilibilil.cf/getdnslist")
	if e != nil {
		return nil, e
	}
	dns := []string{}
	e = json.Unmarshal(bytes, &dns)
	if e != nil {
		return nil, e
	}
	return dns, nil
}

// 负载均衡Dns
//func LoadDns() string {
//
//}
