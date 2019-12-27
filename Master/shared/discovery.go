/**
 * @Author: DollarKillerX
 * @Description: discovery 相关的共享内存
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:50 2019/12/27
 */
package shared

import (
	"ddbf/Master/definition"
	"sync"
)

// 服务注册后的存在空间
var ServerPool = make([]*definition.Server, 0)
var ServerPoolRw sync.RWMutex
