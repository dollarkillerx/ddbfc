/**
 * @Author: DollarKillerX
 * @Description: model.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:24 2019/11/26
 */
package model

import (
	"ddbf/utils/queue"
	"ddbf/utils/set"
)

var BaseModel *baseModel

type baseModel struct {
	Domain  string  // 域名
	Dic     set.Set // 字典
	TimeOut int     // 查询超时
	TryNum  int     // 尝试次数
	Max     int     // 最大并发数量
	OutFile string  // 输出文件
	Death   bool    // 基础扫描 or 死磕到底

	DomainQueue *queue.Queue
	DomainEnd   chan bool
}

func init() {
	BaseModel = &baseModel{
		TimeOut:     300,
		TryNum:      3,
		Max:         200,
		DomainQueue: &queue.Queue{},
		DomainEnd:   make(chan bool, 0),
	}
}
