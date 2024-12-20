package main

import (
	"fmt"
	hello "github.com/Lu271/rpc-test/hello-server/kitex_gen/hello/helloservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	addr, _ := net.ResolveTCPAddr("tcp", ":8090")

	svr := hello.NewServer(
		new(HelloServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		fmt.Println(err)
	}
}
