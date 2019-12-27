/**
 * @Author: DollarKillerX
 * @Description: dns查询连接池
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:57 2019/11/27
 */
package utils

import (
	"context"
	"ddbf/Work/model"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"github.com/miekg/dns"
)

var dnsList = []string{
	"1.1.1.1",
	"1.0.0.1",
	"8.8.8.8",
	"8.8.4.4",
	"9.9.9.9",
	"9.9.9.10",
	"208.67.222.222",
	"208.67.220.220",

	"77.88.8.8",
	"77.88.8.1",
	"156.154.70.22",
	"156.154.71.22",
	"216.146.36.36",
	"216.146.35.35",
	"4.2.2.2",
	"4.2.2.1",
}

// 高效dns连接池

// 单个dns连接
type DnsCoon struct {
	dns     *dns.Conn // dns连接
	version int       // 当前连接的版本号
	host    string    // 当前连接那个dns
}

var NoDomain = errors.New("NoDomain")
var TimeOut = errors.New("TimeOut")

// 解析域名  是否存在,本次查询是否出错
func (d *DnsCoon) DnsParse(ctx context.Context, domain string) (string, error) {
	msg := &dns.Msg{}
	// 创建一个查询消息体
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	// 与当前dns建立连接
	e := d.dns.SetWriteDeadline(time.Now().Add(time.Millisecond * 200))
	if e != nil {
		return "", e
	}

	// 发送查询数据
	e = d.dns.WriteMsg(msg)
	if e != nil {
		return "", e
	}

	// 上面的发送  下面是接受dns消息
	e = d.dns.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
	if e != nil {
		return "", e
	}

	readMsg, e := d.dns.ReadMsg()
	if e != nil || len(readMsg.Question) == 0 {
		// 如果没有数据
		if e != nil {
			return "", e
		} else {
			return "", errors.New("not data")
		}
	}
	if len(readMsg.Answer) == 0 {
		return "", NoDomain
	} else {
		return d.host, nil
	}
}

// dns连接池
type dnsPool struct {
	bufChan chan *DnsCoon
}

var DnsPool *dnsPool

var dnsNum = 1000 // 设在池容量

// 初始化dns连接池
func init() {
	if model.BaseModel.Max > 2000 {
		cs := model.BaseModel.Max / 2
		if cs > 1000 {
			dnsNum = cs
		}
	}

	// 初始化池
	DnsPool = &dnsPool{bufChan: make(chan *DnsCoon, dnsNum)}
	// 初始化池中数据
	for i := 0; i < dnsNum; i++ {
		conn, s, e := getRandDnsCoon()
		if e != nil {
			clog.PrintWa(e)
			continue
		}
		dns := &DnsCoon{host: s, dns: conn, version: 0}
		DnsPool.bufChan <- dns
	}
}

func GetRandDnsCoon() (*DnsCoon, error) {
	conn, s, e := getRandDnsCoon()
	if e != nil {
		return nil, e
	}
	return &DnsCoon{
		dns:     conn,
		version: 0,
		host:    s,
	}, nil
}

// 获取随机dns连接
func getRandDnsCoon() (*dns.Conn, string, error) {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(len(dnsList))
	host := dnsList[intn] + ":53"
	conn, err := dns.DialTimeout("udp", host, time.Second)
	return conn, host, err
}

// 从dns连接池中获得连接
func GetDnsByPool(timeout time.Duration) (*DnsCoon, error) {
	select {
	case dns := <-DnsPool.bufChan:
		return dns, nil
		//case <-time.After(timeout):
		//	return nil, errors.New("time out")
	}
}

// 放回连接
func ReleaseDns(dns *DnsCoon) error {
	select {
	case DnsPool.bufChan <- dns:
		return nil
	default:
		return errors.New("pool overflow")
	}
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
