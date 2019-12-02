/**
 * @Author: DollarKillerX
 * @Description: 深入dns查询
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:31 2019/12/2
 */
package utils

import (
	"github.com/miekg/dns"
	"log"
	"time"
)

//func DnsParsing2(domain string, timeout int, tryNum int) error {
//	var err error
//	for i := 0; i < tryNum; i++ {
//		timeo := time.Now().Add(time.Second)
//		conn, er := newDns("8.8.8.8:53")
//		if er != nil {
//			continue
//		}
//		// 进行dns查询
//		msg := &dns.Msg{}
//		// 进行A记录的查询
//		msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
//		if err = conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
//			clog.PrintEr(err)
//			continue
//		}
//
//		if err = conn.WriteMsg(msg); err != nil {
//			return TimeOut
//			continue
//		}
//
//		if err = conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
//			clog.PrintEr(err)
//			continue
//		}
//		if msg, err = conn.ReadMsg(); err != nil || len(msg.Question) == 0 {
//			if err != nil {
//				if checkTimeOut(err,) {
//					return TimeOut
//				}
//			}
//			continue
//		}
//
//		if time.Now().After(timeo) || time.Now().Equal(timeo) {
//			return TimeOut
//		}
//
//		record := NewRecord(domain, msg.Answer)
//		if record != nil {
//			return nil
//		}
//
//	}
//	if err == nil {
//		err = errors.New("dns err")
//	}
//	return err
//}

func DnsParsing2(domain string) {
	conn, e := newDns("8.8.8.8:53")
	if e != nil {
		log.Panic(e)
	}
	defer conn.Close()

	msg := &dns.Msg{}
	// 创建一个查询消息体
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	// 与当前dns建立连接
	e = conn.SetWriteDeadline(time.Now().Add(time.Second))
	if e != nil {
		panic(e)
	}

	// 发送查询数据
	e = conn.WriteMsg(msg)
	if e != nil {
		panic(e)
	}

	// 以下是接受数据阶段
	readMsg, e := conn.ReadMsg()
	if e != nil || len(readMsg.Question) == 0 {
		// 如果没有数据
		log.Panic("not data")
	}
	log.Println("Answer ", readMsg.Answer)
	log.Println("Ns ", readMsg.Ns)

	// 这部就告诉你bbb就可以了
	//log.Println("in")
	//for _, item := range readMsg.Answer {
	//	//log.Println(item.Header().Name)
	//	//log.Println(item.Header().Header())
	//	//log.Println()
	//	fmt.Println(item.String())
	//	fmt.Println(item.Header().)
	//}
}
