/**
 * @Author: DollarKillerX
 * @Description: engine 总调度器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:51 2019/11/26
 */
package cmd

import (
	"context"
	"ddbf/Work/model"
	"ddbf/Work/utils"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * 优化思路 减少锁的开销
 */

type Engine struct {
}

func EngineInit() {
	engine := Engine{}
	engine.Run()
}

//var ac int64

func (e *Engine) Run() {
	e.initDic() // 初始化字典
	log.Println("[200 OK] 字典初始化完毕")
	log.Println("[200 OK] 进入暴力破解周期")
	log.Println("当前系统并发数: ", model.BaseModel.Max)
	log.Println("当前系统尝试次数: ", model.BaseModel.TryNum)
	e.start() // 开启爆破任务
}

// 开启爆破任务
func (e *Engine) start() {
	t := time.Now()                                   // 计时器
	len := model.BaseModel.Dic.Len()                  // 本次执行的消息总数
	bus := make(chan string, len/10)                  // 任务总线
	out := make(chan []string, model.BaseModel.Max*2) // 输出

	wg := &sync.WaitGroup{}
	wg.Add(model.BaseModel.Max + 2)

	// 初始化chan中的数据
	go e.initChan(wg, bus)

	// 打印日志
	go e.printLog(wg, t, len, bus, out)

	// 暴力破解
	for i := 0; i < model.BaseModel.Max; i++ {
		go e.task(wg, bus, out)
	}

	wg.Wait()
}

type DnsResult struct {
	Domain string // 域名
	Dns    string // 解析它成功的dns
	Ips    []net.IP
}

var jsq uint64

func (e *Engine) task(wg *sync.WaitGroup, bug chan string, out chan []string) {
	defer func() {
		wg.Done()
	}()
loop:
	for {
		select {
		case domain, ok := <-bug:
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
					switch model.BaseModel.Death {
					case true:
						if err.Error() == "dns: bad rdata" || err == utils.NoDomain {
							// 如果这个域名是没有效果的
							atomic.AddUint64(&jsq, 1)
							continue
						}
					default:
						if err.Error() == "dns: bad rdata" || err == utils.NoDomain || err == utils.TimeOut {
							// 如果这个域名是没有效果的
							atomic.AddUint64(&jsq, 1)
							continue
						}
					}
					bug <- domain
					continue
				}

				atomic.AddUint64(&jsq, 1)
				// 如果这个域名是有效的
				// 如果可行 写入到domain中
				re := []string{domain, dnsHost}
				out <- re
			} else {
				break loop
			}
		}
	}
}

// 初始化chan
func (e *Engine) initChan(wg *sync.WaitGroup, bus chan string) {
	defer wg.Done()
	bus <- model.BaseModel.Domain
	for k := range model.BaseModel.Dic {
		domain := strings.TrimSpace(k) + "." + strings.TrimSpace(model.BaseModel.Domain)
		bus <- domain
	}
}

// 打印日志
func (e *Engine) printLog(wg *sync.WaitGroup, tic time.Time, lens int, bus chan string, out chan []string) {
	defer wg.Done()

	one := time.NewTicker(time.Second)
	wu := time.NewTicker(time.Second * 5)
	go func() {
		file, err := os.Create(model.BaseModel.OutFile)
		if err != nil {
			log.Println("文件创建失败")
			log.Println(model.BaseModel.OutFile)
			panic(err)
		}
		defer file.Close()
	loop:
		for {
			select {
			case <-one.C:
				val := atomic.LoadUint64(&jsq)
				if int(val) >= lens+1 {
					log.Println("===============================", val)
					// 程序完结
					time.Sleep(time.Millisecond * 200)
					log.Println(">>>>>>>>>>>>程序完结<<<<<<<<<<<<<<")
					log.Println("字典总长度: ", lens)
					log.Println("总耗时: ", time.Since(tic))
					log.Println(">>>>>>>>>>>>程序完结End<<<<<<<<<<<<<<")
					close(model.BaseModel.DomainEnd)
				}
			case <-wu.C:
				val := atomic.LoadUint64(&jsq)
				fmt.Println("=======================")
				fmt.Println("已完成任务数: ", val)
				fmt.Println("总任务数: ", lens)
				fmt.Println("=======================")
			case <-model.BaseModel.DomainEnd:
				close(bus)
				close(out)
				break loop
			case domain := <-out:
				dom := fmt.Sprintf("成功: %v\n", domain)
				_, err := file.WriteString(dom)
				if err != nil {
					log.Println("文件写入失败")
				}
			}
		}
	}()
}
