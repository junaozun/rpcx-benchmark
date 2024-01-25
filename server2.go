package main

import (
	"github.com/junaozun/rpcx-benchmark/server_func"
	"github.com/junaozun/rpcx-benchmark/util"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/libkv/store"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	_ "net/http/pprof"
	"time"
)

var etcdServer2 *server.Server

const serverAddr2 = "localhost:4835"

func NewEtcdv3Server2(serverAddr, etcdAddr string) (srv *server.Server, err error) {
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
	err = r.Start()
	if err != nil {
		return
	}
	srv.Plugins.Add(r)
	return
}

func init() {
	srv, err := NewEtcdv3Server2(serverAddr2, "http://127.0.0.1:2379")
	if err != nil {
		return
	}
	etcdServer2 = srv
}

func main() {
	//go func() {
	//	http.ListenAndServe("0.0.0.0:8899", http.DefaultServeMux)
	//}()
	err := etcdServer2.RegisterName("HelloTest", new(server_func.GreeterImpl), "httpnfo=dddjejadjflds")
	if err != nil {
		return
	}
	err = etcdServer2.RegisterName("GvgTest", new(server_func.GvgTestImpl), "httpnfo=dddjejadjflds")
	if err != nil {
		return
	}
	err = etcdServer2.Serve("tcp", serverAddr2)
	if err != nil {
		return
	}
}
