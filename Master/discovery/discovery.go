/**
 * @Author: DollarKillerX
 * @Description: 服务发现
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午12:06 2019/12/27
 */
package discovery

import (
	"context"
	"ddbf/Master/definition"
	"ddbf/Master/shared"
	"ddbf/pb/pb_master"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

type discovery struct {
}

type report struct {
}

func (d *discovery) Register(ctx context.Context, request *pb_master.DiscoveryRequest) (*pb_master.DiscoveryResponse, error) {
	var server definition.Server
	// 如果服务id为空说明改服务还未被注册
	if request.WorkId == "" {
		server = definition.Server{}
		workId := easyutils.SuperRand()
		server.Id = workId
		server.Host = request.Host
		server.Load = request.Load
		server.TimeOut = time.Now().Add(time.Millisecond * 400) // 设置心跳检测时间

		// 注册服务
		shared.ServerPoolRw.RLock() // 这里为什么写却加读锁  看下心跳服务你就知道了
		shared.ServerPool = append(shared.ServerPool, &server)
		shared.ServerPoolRw.RUnlock()
		log.Printf("服务注册成功  ServerId:%s ServerHost:%s \n", workId, request.Host)
		return &pb_master.DiscoveryResponse{
			WorkId: workId,
		}, nil
	}
	// 应用续心跳
	// 这里要考虑到一个问题  如果发生网络波动或者脑裂  (服务是存在的但是当时网络不通)  [丢弃当前包,让对方重启]  (golang可以实现热重启)
	// 上面那方法有点麻烦 我就把她 添加到新的上面去吧
	foundIt := false
	for _, v := range shared.ServerPool {
		if v.Id == request.WorkId {
			foundIt = true
			v.Load = request.Load
			v.TimeOut = time.Now().Add(time.Millisecond * 400)
			v.TryNum = 0
			break
		}
	}
	if !foundIt {
		// 如果哈哈哈 就重新注册
		server = definition.Server{}
		server.Load = request.Load
		server.Id = request.WorkId
		server.Host = request.Host
		server.TimeOut = time.Now().Add(time.Millisecond * 400)

		// 注册服务
		shared.ServerPoolRw.RLock()
		shared.ServerPool = append(shared.ServerPool, &server)
		shared.ServerPoolRw.RUnlock()
	}

	return &pb_master.DiscoveryResponse{
		WorkId: request.WorkId,
	}, nil
}

func (r *report) Report(ctx context.Context, request *pb_master.TaskReport) (*pb_master.TaskResponse, error) {
	log.Printf("任务入库200OK TaskId: %s", request.TaskId)
	_, ex := shared.TaskRun[request.WorkId]
	if !ex {
		// 如果不存在丢弃这个数据包
		return &pb_master.TaskResponse{
			StatusCode: 200,
		}, nil
	}

	task := shared.TaskNum[request.TaskId]
	task.Lock()
	defer task.Unlock()
	task.OverNum++
	file, e := getFile(request.TaskId)
	if e != nil {
		clog.PrintWa(e)
		return &pb_master.TaskResponse{
			StatusCode: 500,
		}, nil
	}
	defer file.Close()
	for _, v := range request.TaskItem {
		file.WriteString(v.Domain + "  :  " + v.DnsHost + "\n")
	}

	if task.Over {
		if task.OverNum >= task.Num {
			<-shared.LimitChannel
			file.WriteString("Over 处理完毕 \n")
		}
	}

	return &pb_master.TaskResponse{
		StatusCode: 200,
	}, nil
}

func RunServer(host string) {
	listener, e := net.Listen("tcp", host)
	if e != nil {
		log.Fatalln(e)
	}
	server := grpc.NewServer()
	pb_master.RegisterReportServer(server, &report{})
	pb_master.RegisterRegisteredWorkServer(server, &discovery{})
	go heartbeatCheck() // 心跳检测
	if e = server.Serve(listener); e != nil {
		log.Fatalln(e)
	}
}

// 心跳检测服务
func heartbeatCheck() {
	ticker := time.NewTicker(time.Millisecond * 250)
	for {
		select {
		case <-ticker.C:
			for _, v := range shared.ServerPool {
				// 如果发生超时
				if v.TimeOut.Before(time.Now()) {
					log.Printf("应用 [ %s : %s ] 发生网络抖动", v.Id, v.Host)
					if v.TryNum < 3 {
						v.TryNum++
					} else {
						log.Printf("应用 [ %s : %s ] 以掉线", v.Id, v.Host)

						shared.ServerPoolRw.Lock()
						for k, vv := range shared.ServerPool {
							if vv.Id == v.Id {
								// 删除这一个
								if k != len(shared.ServerPool) {
									shared.ServerPool = append(shared.ServerPool[:k], shared.ServerPool[k+1:]...)
								} else {
									shared.ServerPool = shared.ServerPool[:k]
								}
								break
							}
						}
						shared.ServerPoolRw.Unlock()

						// 如果这个超时服务 有任务  就把这个任务放回到原来的任务队列中
						task, ex := shared.TaskRun[v.Id]
						if ex {
							shared.TaskPool <- *task.Data
							delete(shared.TaskRun, v.Id)
						}
					}
				} else {
					v.TryNum = 0 // 如果没有超时 清空他的超时记录
				}
			}
		}
	}
}

// 如果日志文件不存在就创建日志文件
func getFile(id string) (*os.File, error) {
	filepath := filepath.Join(definition.OUTFILE, id+".txt")
	return os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 00666)
}
