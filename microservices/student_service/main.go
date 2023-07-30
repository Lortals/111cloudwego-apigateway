package main

import (
	"log"
	"net"

	studentservice "github.com/Lortals/111cloudwego-apigateway/microservices/student_service/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := studentservice.NewServer(new(StudentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "student"}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
