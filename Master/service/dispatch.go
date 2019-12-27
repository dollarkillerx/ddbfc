/**
 * @Author: DollarKillerX
 * @Description: dispatch 消息调度
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:48 2019/12/27
 */
package service

import (
	"context"
	"ddbf/Master/definition"
	"ddbf/Master/grpc_conn"
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
					blast, e := grpc_conn.BlastNew(k.Host)
					if e != nil {
						log.Println(e)
						continue
					}
					response, e := blast.Task(context.TODO(), &data)
					if e != nil {
						log.Println(e)
						continue
					}

					if response.StatusCode != 200 {
						continue
					}
					// 进入这里的话 就说明正常发布了
					// 开始记录状态
					shared.TaskRun[k.Id] = &definition.TaskItemR{Data: &data}
					vol = true
					break
				}
			}
			if !vol {
				shared.TaskPool <- data
			}

		}
	}
}
