/**
 * @Author: DollarKillerX
 * @Description: grpc server的实现
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:19 2019/12/27
 */
package grpc_server

import (
	"context"
	"ddbf/Work/service"
	"ddbf/Work/shared"
	"ddbf/pb/pb_work"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Task(ctx context.Context, req *pb_work.Request) (*pb_work.Response, error) {
	shared.LoadRw.RLock()
	ac := shared.Load
	shared.LoadRw.Unlock()
	if ac != 0 {
		return &pb_work.Response{
			StatusCode: 2503,
		}, nil
	}

	// 如果服务还没有负载
	shared.LoadRw.Lock()
	shared.Load++
	shared.LoadRw.Unlock()

	// 接入数据
	shared.TaskId = req.TaskId

	// 注入数据
	go service.InitWorkDispatch(req.TaskItem)

	return &pb_work.Response{
		StatusCode: 200,
	}, nil
}

//func insertData(items []string) {
//	shared.TaskNum = len(items)
//	for _, v := range items {
//		shared.TaskChannel <- v
//	}
//}

func RunServer(host string) {
	listener, e := net.Listen("tcp", host)
	if e != nil {
		log.Fatalln(e)
	}
	ser := grpc.NewServer()
	pb_work.RegisterTaskServer(ser, &server{})
	if e = ser.Serve(listener); e != nil {
		log.Fatalln(e)
	}
}
