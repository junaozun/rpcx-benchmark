package server_func

import (
	"context"
	"github.com/junaozun/rpcx-benchmark/pb"
)

type GreeterImpl struct{}

func (s *GreeterImpl) SayHello(ctx context.Context, args *pb.HelloRequest, reply *pb.HelloResponse) (err error) {
	*reply = pb.HelloResponse{
		Result: "suxuefeng",
		Header: map[string]string{
			"naifdhaof":   "sdfjalsdfjkwe",
			"38459324":    "sdfjalsdfjkwe",
			"8345245jhfj": "sdfjalsdfjkwe",
		},
		PageInfo: map[string]int64{
			"skajdfdf444": 12,
			"77656":       98,
			"jfksadj":     1872,
		},
		TType:  "categoru",
		Status: "active",
		Code:   "13923",
	}
	return nil
}

type GvgTestImpl struct{}

func (s *GvgTestImpl) GetMapInfo(ctx context.Context, args *pb.HelloRequest, reply *pb.HelloResponse) (err error) {
	*reply = pb.HelloResponse{
		Result: "suxuefeng",
		Header: map[string]string{
			"naifdhaof":   "sdfjalsdfjkwe",
			"38459324":    "sdfjalsdfjkwe",
			"8345245jhfj": "sdfjalsdfjkwe",
		},
		PageInfo: map[string]int64{
			"skajdfdf444": 12,
			"77656":       98,
			"jfksadj":     1872,
		},
		TType:  "categoru",
		Status: "active",
		Code:   "13923",
	}
	return nil
}
