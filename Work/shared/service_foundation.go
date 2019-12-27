/**
 * @Author: DollarKillerX
 * @Description: 服务的基础定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:10 2019/12/27
 */
package shared

import "sync"

// 服务的基础定义
var (
	WorkId = "" // 当前服务的id

	LoadRw sync.Mutex
	Load   = 0 // 当前服务的负载

	TaskId = "" // 任务的ID
)
