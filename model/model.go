/**
 * @Author: DollarKillerX
 * @Description: model.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:24 2019/11/26
 */
package model

import (
	"time"

	"ddbf/utils/set"
)

var BaseModel *baseModel

type baseModel struct {
	Domain  string        // 域名
	Dic     set.Set       // 字典
	TimeOut time.Duration // 查询超时
	TryNum  int           // 尝试次数
}

func init() {
	BaseModel = &baseModel{
		TimeOut: 200 * time.Millisecond,
		TryNum:  3,
	}
}