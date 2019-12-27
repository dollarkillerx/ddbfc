/**
 * @Author: DollarKillerX
 * @Description: Domain Blast相关的连接
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:28 2019/12/27
 */
package grpc_conn

import (
	"context"
	"ddbf/pb/pb_work"
	"google.golang.org/grpc"
)

type blast struct {
	host  string
	blast pb_work.TaskClient
}

func BlastNew(host string) (*blast, error) {
	conn, e := grpc.Dial(host, grpc.WithInsecure())
	if e != nil {
		return nil, e
	}
	client := pb_work.NewTaskClient(conn)
	return &blast{
		host:  host,
		blast: client,
	}, nil
}

func (c *blast) Task(ctx context.Context, req *pb_work.Request) (*pb_work.Response, error) {
	return c.blast.Task(ctx, req)
}
