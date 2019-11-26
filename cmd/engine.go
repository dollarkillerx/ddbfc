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
