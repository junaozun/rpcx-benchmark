package main

import (
	"context"
	"github.com/junaozun/rpcx-benchmark/pb"
	"github.com/junaozun/rpcx-benchmark/util"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/libkv/store"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var etcdServer *server.Server

func NewEtcdv3Server(serverAddr, etcdAddr string) (srv *server.Server, err error) {
	srv = server.NewServer()
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + serverAddr,
		EtcdServers:    []string{etcdAddr},
		BasePath:       "/rpcx_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
		Options:        &store.Config{},
	}
	share.Codecs[protocol.SerializeType(6)] = &util.JsoniterCodec{}
	share.Codecs[protocol.SerializeType(7)] = &util.SonicJson{}
	err = r.Start()
	if err != nil {
		return
	}
	srv.Plugins.Add(r)
	return
}

func init() {
	srv, err := NewEtcdv3Server("localhost:4834", "http://127.0.0.1:2379")
	if err != nil {
		return
	}
	etcdServer = srv
}

func main() {
	err := etcdServer.RegisterName("HelloTest", new(GreeterImpl), "httpnfo=dddjejadjflds")
	if err != nil {
		return
	}
	pprof()
	err = etcdServer.Serve("tcp", "localhost:4834")
	if err != nil {
		return
	}
}

func pprof() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", http.DefaultServeMux)
	}()
}

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
