/**
 * @Author: DollarKillerX
 * @Description: 基于公共数据源的域名发现
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:37 2019/12/4
 */
package server

import (
	"ddbf/datasource"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dollarkillerx/easyutils/clog"
)

// 设计订阅发布模型
type discovery struct {
	dataSource []datasource.BaseInterface // 所以数据源
}

func DiscoveryNew() *discovery {
	return &discovery{}
}

// 总调度器
func (d *discovery) Run(domain string) {
	d.init()
	wg := sync.WaitGroup{}
	bus := make(chan datasource.Domains, 1000)
	// 一个总控
	for _, fun := range d.dataSource {
		wg.Add(1)
		go d.parserDomain(domain, fun, bus, &wg)
	}
	// 一个消息处理的总线 && 数据清洗
	wg.Add(1)
	go d.bus(bus, &wg)

	wg.Wait()
}

// 初始化所有数据源
func (d *discovery) init() {
	d.dataSource = datasource.DataSources
}

// 每个数据源的引擎
func (d *discovery) parserDomain(domain string, pareFunc datasource.BaseInterface, bus chan datasource.Domains, wg *sync.WaitGroup) {
	defer wg.Done()
	domains, e := pareFunc.ParseDomain(domain)
	if e != nil {
		clog.PrintWa(e)
		//bus <- domains
		//return
	}
	bus <- domains
}

// 即使关闭与数据清洗
func (d *discovery) bus(busChan chan datasource.Domains, wg *sync.WaitGroup) {
	defer wg.Done()
	file, e := os.Create("outDns.txt")
	if e != nil {
		log.Fatalln("输出文件创建失败!!!")
	}
	count := 0                     // 计数器
	total := len(d.dataSource) - 1 // 总次数
	db := datasource.Domains{}
	for {
		select {
		case domain := <-busChan:
			count++
			for k, v := range domain {
				_, ok := db[k]
				if !ok {
					db[k] = v
					dom := fmt.Sprintf("成功: %v\n", k)
					file.WriteString(dom)
					//fmt.Println(k, "  ", v)
				}
			}

			if count >= total {
				log.Println("任务执行完毕")
				return
			}
		}
	}
}
