/**
 * @Author: DollarKillerX
 * @Description: engine 核心调度
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:54 2019/12/27
 */
package service

import (
	"context"
	"ddbf/Work/shared"
	"ddbf/Work/utils"
	"ddbf/pb/pb_master"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func InitWorkDispatch(domains []string) {
	total := len(domains)
	bus := make(chan string, 10000)
	out := make(chan *pb_master.DomainItem, 5000)

	wg := &sync.WaitGroup{}
	wg.Add(1002)

	go initChannel(wg, bus, domains)

	go checkOver(wg, total, bus, out)

	// 暴力破解
	for i := 0; i < 1000; i++ {
		go task(wg, bus, out)
	}

	wg.Wait()
}

// 初始化channel
func initChannel(wg *sync.WaitGroup, bus chan string, data []string) {
	defer wg.Done()
	for _, v := range data {
		bus <- v
	}
}

var jsq uint64

func task(wg *sync.WaitGroup, bus chan string, out chan *pb_master.DomainItem) {
	defer wg.Done()
loop:
	for {
		select {
		case domain, ok := <-bus:
			if ok {
				// 向资源池中获取
				pool, err := utils.GetDnsByPool(time.Second * 3)
				if err != nil {
					log.Panic(err)
				}
				timeout, _ := context.WithTimeout(context.TODO(), time.Millisecond*200)

				dnsHost, err := pool.DnsParse(timeout, domain)
				//使用完后放回
				if err := utils.ReleaseDns(pool); err != nil {
					panic(err)
				}

				if err != nil {
					// 如果本次查询错误
					if err.Error() == "dns: bad rdata" || err == utils.NoDomain || err == utils.TimeOut {
						// 如果这个域名是没有效果的
						atomic.AddUint64(&jsq, 1)
						continue loop
					}
					bus <- domain
					continue loop
				}

				atomic.AddUint64(&jsq, 1)
				// 如果这个域名是有效的
				// 如果可行 写入到domain中
				out <- &pb_master.DomainItem{Domain: domain, DnsHost: dnsHost}
			} else {
				break loop
			}
		}
	}
}

func checkOver(wg *sync.WaitGroup, total int, bus chan string, out chan *pb_master.DomainItem) {
	defer wg.Done()
	// 1.监听任务是否完成
	// 2.收集over数据
	// 3.完成后退出所有采集相关
	group := sync.WaitGroup{}
	over := make(chan bool, 1) // 传递退出信号
	ticker500 := time.NewTicker(time.Millisecond * 500)

	group.Add(2)
	// 监听任务是否完毕
	go func() {
		defer group.Done()
		for {
			select {
			case <-ticker500.C:
				num := int(atomic.LoadUint64(&jsq))
				if num >= total {
					// 任务完毕
					clog.PrintWa(num)
					over <- true
					return
				}
			}
		}
	}()

	go func() {
		defer group.Done()
		result := make([]*pb_master.DomainItem, 0)
		for {
			select {
			case data := <-out:
				result = append(result, data)
			case <-over:
				//close(out)
				close(bus)
				shared.Over <- result
				jsq = 0
				return
			}
		}
	}()

	group.Wait()
}
