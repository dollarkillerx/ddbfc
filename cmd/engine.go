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
	log.Println("[200 OK] 字典初始化完毕")
	log.Println("[200 OK] 进入暴力破解周期")
	log.Println("当前系统并发数: ", model.BaseModel.Max)
	log.Println("当前系统尝试次数: ", model.BaseModel.TryNum)
	e.start() // 开启爆破任务
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

// 一些计数器
var mcmu sync.Mutex
var mc = 0
var okmu sync.Mutex
var oktotal = 0

// 开启爆破任务
func (e *Engine) start() {
	t := time.Now() // 计时器
	len := model.BaseModel.Dic.Len()
	bus := make(chan string, len)
	wg := &sync.WaitGroup{}
	wg.Add(model.BaseModel.Max + 2)

	// 初始化chan中的数据
	go e.initChan(wg, bus)

	// 打印日志
	go e.printLog(wg, t, len, bus)

	// 暴力破解
	for i := 0; i < model.BaseModel.Max; i++ {
		go e.task(wg, bus)
	}

	wg.Wait()
}

func (e *Engine) task(wg *sync.WaitGroup, bug chan string) {
	defer func() {
		wg.Done()

		// 这个告诉任务
		mcmu.Lock()
		mc++
		mcmu.Unlock()
	}()
	for {
		select {
		case domain, ok := <-bug:
			if ok {
				err := utils.DnsParsing2(domain, model.BaseModel.TimeOut, model.BaseModel.TryNum)
				if err != nil {
					// 如果超时就回写
					if err == utils.TimeOut {
						bug <- domain
						continue
					}
					okmu.Lock()
					oktotal++
					okmu.Unlock()
					continue
				}
				okmu.Lock()
				oktotal++
				okmu.Unlock()
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
	//close(bus)
}

// 打印日志
func (e *Engine) printLog(wg *sync.WaitGroup, tic time.Time, len int, bus chan string) {
	defer wg.Done()
	go func() {
		for {
			select {
			case <-time.After(time.Second * 10):
				log.Println("===========================")
				if oktotal >= len {
					close(bus)
				}
				okmu.Lock()
				log.Println("已完成任务: ", oktotal)
				okmu.Unlock()
				log.Println("总任务数: ", len)
				log.Println("===========================")
			}
		}
	}()
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				mcmu.Lock()
				if mc >= model.BaseModel.Max {
					// 程序完结
					time.Sleep(time.Millisecond * 200)
					log.Println(">>>>>>>>>>>>程序完结<<<<<<<<<<<<<<")
					log.Println("字典总长度: ", len)
					log.Println("总耗时: ", time.Since(tic))
					log.Println(">>>>>>>>>>>>程序完结End<<<<<<<<<<<<<<")
					close(model.BaseModel.DomainEnd)
				}
				mcmu.Unlock()
			}
		}
	}()
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
