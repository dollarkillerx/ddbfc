/**
 * @Author: DollarKillerX
 * @Description: model.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:24 2019/11/26
 */
package model

import (
	"ddbf/utils/queue"
	"time"

	"ddbf/utils/set"
)

var BaseModel *baseModel

type baseModel struct {
	Domain  string        // 域名
	Dic     set.Set       // 字典
	TimeOut time.Duration // 查询超时
	TryNum  int           // 尝试次数
	Max     int           // 最大并发数量

	DomainQueue *queue.Queue
	DomainEnd   chan bool
}

func init() {
	BaseModel = &baseModel{
		TimeOut:     1000 * time.Millisecond,
		TryNum:      2,
		Max:         200,
		DomainQueue: &queue.Queue{},
		DomainEnd:   make(chan bool, 0),
	}
}
