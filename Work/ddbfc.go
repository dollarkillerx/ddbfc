/**
 * @Author: DollarKillerX
 * @Description: main程序入口
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:17 2019/11/26
 */
package main

import (
	"context"
	"ddbf/Work/discovery"
	"ddbf/Work/grpc_server"
	"ddbf/Work/service"
	"ddbf/Work/shared"
	"ddbf/Work/utils"
	"ddbf/pb/pb_master"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// 进行性能分析
	//go scp()
	if len(os.Args) != 3 {
		log.Fatalln("参数不对!  ./main master地址 server地址")
	}
	host := os.Args[1]                   // 服务注册地址
	serHost := utils.DockerGetHostName() // rpc server地址

	go grpc_server.RunServer(serHost) // 启动grpc服务
	go heartbeat(host, serHost)       // 启动心跳服务

	// 监听的是结束信号
	dis, _ := discovery.DiscoveryNew(host)
	for {
		select {
		case work := <-shared.Over:
			log.Printf("任务处理完毕: %s\n", shared.TaskId)
			result := &pb_master.TaskReport{TaskId: shared.TaskId, WorkId: shared.WorkId}
			resultItems := make([]*pb_master.DomainItem, 0)
			for _, v := range work {
				resultItem := &pb_master.DomainItem{DnsHost: v.Dns, Domain: v.Domain}
				resultItems = append(resultItems, resultItem)
			}

			for {
				_, e := dis.Report(context.TODO(), result)
				if e == nil {
					break
				}
			}
			// 清除状态锁
			shared.LoadRw.Lock()
			shared.Load = 0
			shared.LoadRw.Unlock()
		}
	}
}

// 心跳服务
func heartbeat(host, rpc string) {
	dis, e := discovery.DiscoveryNew(host)
	if e != nil {
		log.Fatalln(e)
	}

	data := &pb_master.DiscoveryRequest{Host: rpc}
	response, e := dis.Register(context.TODO(), data)
	if e != nil {
		log.Fatalln(e)
	}
	shared.WorkId = response.WorkId
	data.WorkId = response.WorkId

	ticker := time.NewTicker(time.Millisecond * 200)
	for {
		select {
		case <-ticker.C:

			shared.LoadRw.RLock()
			data.Load = int64(shared.Load)
			shared.LoadRw.RUnlock()

			_, e := dis.Register(context.TODO(), data)
			if e != nil {
				log.Println("心跳服务 Err ", e)
			}
		}
	}
}

func scp() {
	cpuf, e := os.Create("cpu_profile")
	if e != nil {
		log.Fatalln(e)
	}
	pprof.StartCPUProfile(cpuf)

	time.Sleep(60 * time.Second)
	defer pprof.StopCPUProfile()

	log.Println("关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||  CPU")

	memf, err := os.Create("mem_profile")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	memf.Close()

	log.Println("关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析||| MEMF")
}
