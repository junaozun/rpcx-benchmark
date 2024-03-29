// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// source: hello.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	client "github.com/smallnest/rpcx/client"
	server "github.com/smallnest/rpcx/server"
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

var _HelloTest_serviceDesc = &server.ServiceDesc{
	ServiceName: "HelloTest",
	Methods:     []server.MethodDesc{
		{
			MethodName:  "SayHello",
			Handler:     _HelloTest_Handler,
		},
	},
	HandlerType: (*IHelloTestService)(nil),
}

// =================== IHelloTestService ===================
type IHelloTestService interface {
	SayHello(ctx context.Context, req *HelloRequest, reply *HelloResponse) (err error)
}

func RegisterHelloTestService(s *server.Server, svr interface{},metadata string)error {
	return s.RegisterService(_HelloTest_serviceDesc, svr,metadata)
}

func _HelloTest_Handler(svc interface{},ctx context.Context,req interface{},reply any) error {
	return svc.(IHelloTestService).SayHello(ctx,req.(*HelloRequest),reply.(*HelloResponse))
}

// =================== IHelloTestEtcdV3Client ===================
type IHelloTestEtcdV3Client interface {
	SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error)
}

type HelloTestEtcdV3Client struct {
	xclient client.XClient
}

func NewHelloTestEtcdV3Client(xclient client.XClient) IHelloTestEtcdV3Client {
	ret := &HelloTestEtcdV3Client{xclient: xclient}
	return ret
}
func (c *HelloTestEtcdV3Client) SayHello(ctx context.Context, req *HelloRequest) (reply *HelloResponse, err error) {
	reply = &HelloResponse{}
	err = c.xclient.Call(ctx, "SayHello", req, reply)
	return reply, err
}

// =================== IGvgTestService ===================
type IGvgTestService interface {
	GetMapInfo(ctx context.Context, req *HelloRequest, reply *HelloResponse) (err error)
}

// =================== IGvgTestEtcdV3Client ===================
type IGvgTestEtcdV3Client interface {
	GetMapInfo(ctx context.Context, req *HelloRequest) (*HelloResponse, error)
}

type GvgTestEtcdV3Client struct {
	xclient client.XClient
}

func NewGvgTestEtcdV3Client(xclient client.XClient) IGvgTestEtcdV3Client {
	ret := &GvgTestEtcdV3Client{xclient: xclient}
	return ret
}
func (c *GvgTestEtcdV3Client) GetMapInfo(ctx context.Context, req *HelloRequest) (reply *HelloResponse, err error) {
	reply = &HelloResponse{}
	err = c.xclient.Call(ctx, "GetMapInfo", req, reply)
	return reply, err
}
