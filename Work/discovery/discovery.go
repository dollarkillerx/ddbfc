/**
 * @Author: DollarKillerX
 * @Description: discovery 服务发现相关  (心跳)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:09 2019/12/27
 */
package discovery

import (
	"context"
	"ddbf/pb/pb_master"
	"google.golang.org/grpc"
)

type discovery struct {
	host     string
	register pb_master.RegisteredWorkClient
	report   pb_master.ReportClient
}

func DiscoveryNew(host string) (*discovery, error) {
	conn, e := grpc.Dial(host, grpc.WithInsecure())
	if e != nil {
		return nil, e
	}
	report := pb_master.NewReportClient(conn)
	register := pb_master.NewRegisteredWorkClient(conn)

	return &discovery{
		host:     host,
		report:   report,
		register: register,
	}, nil
}

func (d *discovery) Register(ctx context.Context, req *pb_master.DiscoveryRequest) (*pb_master.DiscoveryResponse, error) {
	return d.register.Register(ctx, req)
}

func (d *discovery) Report(ctx context.Context, req *pb_master.TaskReport) (*pb_master.TaskResponse, error) {
	return d.report.Report(ctx, req)
}
