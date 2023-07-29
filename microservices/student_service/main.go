package main

import (
	"log"
	"net"

	management "github.com/ararchch/api-gateway/microservices/student_service/kitex_gen/student/api/studentmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	// initate new etcd registry at port 2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        log.Fatal(err)
    }

	// create new Kitex server for Student Service
	svr := management.NewServer(
		new(StudentManagementImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Student"}),
		server.WithRegistry(r), 
		server.WithServiceAddr(&net.TCPAddr{Port: 8889}),
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
