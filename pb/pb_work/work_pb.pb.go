// Code generated by protoc-gen-go. DO NOT EDIT.
// source: work_pb.proto

package pb_work

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskItem             []string `protobuf:"bytes,3,rep,name=task_item,json=taskItem,proto3" json:"task_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bba6fc18c78a732d, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Request) GetTaskItem() []string {
	if m != nil {
		return m.TaskItem
	}
	return nil
}

type Response struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bba6fc18c78a732d, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "pb_work.Request")
	proto.RegisterType((*Response)(nil), "pb_work.Response")
}

func init() { proto.RegisterFile("work_pb.proto", fileDescriptor_bba6fc18c78a732d) }

var fileDescriptor_bba6fc18c78a732d = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xcf, 0x2f, 0xca,
	0x8e, 0x2f, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0x8a, 0x07, 0x89,
	0x28, 0xd9, 0x73, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x73, 0xb1, 0x97,
	0x24, 0x16, 0x67, 0xc7, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x81, 0xb8,
	0x9e, 0x29, 0x42, 0xd2, 0x5c, 0x9c, 0x10, 0x89, 0x92, 0xd4, 0x5c, 0x09, 0x66, 0x05, 0x66, 0x0d,
	0xce, 0x20, 0x0e, 0xb0, 0x54, 0x49, 0x6a, 0xae, 0x92, 0x36, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x3c, 0x17, 0x77, 0x71, 0x49, 0x62, 0x49, 0x69, 0x71, 0x7c, 0x72,
	0x7e, 0x4a, 0x2a, 0xd8, 0x14, 0xe6, 0x20, 0x2e, 0x88, 0x90, 0x73, 0x7e, 0x4a, 0xaa, 0x91, 0x31,
	0x17, 0x4b, 0x48, 0x62, 0x71, 0xb6, 0x90, 0x36, 0x94, 0x16, 0xd0, 0x83, 0xba, 0x43, 0x0f, 0xea,
	0x08, 0x29, 0x41, 0x24, 0x11, 0x88, 0xa9, 0x49, 0x6c, 0x60, 0x27, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x52, 0xea, 0x79, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	Task(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Task(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb_work.Task/Task", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	Task(context.Context, *Request) (*Response, error)
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) Task(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Task not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_Task_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Task(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_work.Task/Task",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Task(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_work.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Task",
			Handler:    _Task_Task_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "work_pb.proto",
}