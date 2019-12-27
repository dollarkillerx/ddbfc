/**
 * @Author: DollarKillerX
 * @Description: task 相关状态
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:11 2019/12/27
 */
package shared

import (
	"ddbf/Master/definition"
	"ddbf/pb/pb_work"
)

// 统计 每一个大的task下任务的数量
var TaskNum = make(map[string]*definition.Task)

// 任务池
var TaskPool = make(chan pb_work.Request, 100)

// 任务表   那个服务正在执行什么任务
var TaskRun = make(map[string]*definition.TaskItemR) // 服务id ： 服务
