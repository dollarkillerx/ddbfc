/**
 * @Author: DollarKillerX
 * @Description: grpc server的实现
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:19 2019/12/27
 */
package grpc_server

import (
	"context"
	"ddbf/pb/pb_work"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) Task(ctx context.Context, req *pb_work.Request) (*pb_work.Response, error) {

	return &pb_work.Response{}, nil
}

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
