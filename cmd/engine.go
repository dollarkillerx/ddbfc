/**
 * @Author: DollarKillerX
 * @Description: engine 总调度器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:51 2019/11/26
 */
package cmd

import (
	"ddbf/model"
	"ddbf/utils"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

// 相关基础定义
const (
	dicDir  = "dic"                                                                    // 字典文件所在目录
	dicPath = "dic/base.dic"                                                           // 字典文件路径
	dicUlr  = "https://raw.githubusercontent.com/dollarkillerx/ddbfc/cli/dic/base.dic" // 字典文件Url地址
)

type Engine struct {
}

func EngineInit() {
	engine := Engine{}
	engine.Run()
}

func (e *Engine) Run() {
	e.initDic() // 初始化字典
	e.start()   // 开启爆破任务
}

// 初始化字典
func (e *Engine) initDic() {
	// 检测dic目录 是否为空,如果为空 使用默认字典
	empty := utils.FileDirEmpty(dicDir)
	if empty {
		// 为空 使用默认字典
		e.defaultDic()
	}
	// 读取dic目录下的字典 写入到 字典set中
	sets, err := utils.LoopDir(dicDir)
	if err != nil {
		log.Fatalln(err)
	}
	model.BaseModel.Dic = sets
}

func (e *Engine) defaultDic() {
	err := utils.DirPing(dicDir)
	if err != nil {
		log.Fatalln("Directory creation failed")
	}
	// 向github下载最新的默认字典
	easyHttp := utils.EasyHttpNew()
	bytes, err := easyHttp.Get(dicUlr)
	if err != nil {
		log.Fatalln("GitHub download default dictionary failed")
	}

	// 默认字典下载完毕 存入其中
	err = ioutil.WriteFile(dicPath, bytes, 000755)
	if err != nil {
		log.Fatalln(err)
	}
}

var mcmu sync.Mutex
var mc = 0

// 开启爆破任务
func (e *Engine) start() {
	len := model.BaseModel.Dic.Len()
	bus := make(chan string, len)
	wg := &sync.WaitGroup{}
	wg.Add(model.BaseModel.Max + 2)

	// 初始化chan中的数据
	go e.initChan(wg, bus)

	// 打印日志
	go e.printLog(wg)

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				mcmu.Lock()
				if mc >= model.BaseModel.Max {
					close(model.BaseModel.DomainEnd)
				}
				mcmu.Unlock()
			}
		}
	}()

	// 暴力破解
	for i := 0; i < model.BaseModel.Max; i++ {
		go e.task(wg, bus)
	}

	wg.Wait()
}

func (e *Engine) task(wg *sync.WaitGroup, bug chan string) {
	defer func() {
		wg.Done()

		mcmu.Lock()
		mc++
		mcmu.Unlock()
	}()
	for {
		select {
		case domain, ok := <-bug:
			if ok {
				err := utils.DnsParsing(domain, model.BaseModel.TimeOut, model.BaseModel.TryNum)
				if err != nil {
					if domain == "www.dollarkiller.com" || domain == "translate.dollarkiller.com" {
						log.Println(err)
					}
					continue
				}
				// 如果可行 写入到domain中
				model.BaseModel.DomainQueue.Append(domain)
			} else {
				return
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
		//log.Println(domain)
		bus <- domain
	}
	close(bus)
}

// 打印日志
func (e *Engine) printLog(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-model.BaseModel.DomainEnd:
			return
		default:
			next, b := model.BaseModel.DomainQueue.Next()
			if b {
				log.Printf("成功: %v", next)
			}
		}
	}
}
