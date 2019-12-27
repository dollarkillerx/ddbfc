/**
 * @Author: DollarKillerX
 * @Description: dispatch 消息调度
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:48 2019/12/27
 */
package service

import (
	"ddbf/Master/shared"
	"log"
)

// 消费服务 把当前的应用发送个给消费者
func Dispatch() {
	for {
		select {
		case data := <-shared.TaskPool:
			log.Printf("[Dispath] Log 当前任务大小: %d", len(data.TaskItem))
			vol := false // 判断当前任务是否被消费
			for _, k := range shared.ServerPool {
				if k.Load == 0 {

				}
			}

		}
	}
}
