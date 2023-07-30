package main

import (
	"log"

	management "github.com/Lortals/111cloudwego-apigateway/microservices/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	// initate new etcd registry at port 2370
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2370"})
	if err != nil {
		log.Fatal(err)
	}

	// create new Kitex server for Multiplication Service
	svr := management.NewServer(
		new(MultiplicationManagementImpl), // Follow MultiplicationManagementImpl as defined in ./handler.go
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Multiplication"}), // allow service to be discovered with name: "Multiplication"
		server.WithRegistry(r), // register service on etcd registry 'r' (as declared earlier)
	)

	// run server and handler any errors
	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
