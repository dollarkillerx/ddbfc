/**
 * @Author: DollarKillerX
 * @Description: task 任务的定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:05 2019/12/27
 */
package definition

import (
	"ddbf/pb/pb_work"
	"sync"
)

// task表
type Task struct {
	Id  string // task id
	Num int    // 这个id下任务的数量

	Over       bool // 是否发送完毕
	sync.Mutex      // 用于写文件时防止错误
	OverNum    int  // 已完成的任务
}

// 单个任务
type TaskItemR struct {
	//sync.Mutex
	Data *pb_work.Request
}
