package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/Lortals/111cloudwego-apigateway/microservices/student_service/kitex_gen/demo"
	studentservice "github.com/Lortals/111cloudwego-apigateway/microservices/student_service/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2370"})
	if err != nil {
		log.Fatal(err)
	}

	cli, err := studentservice.NewClient("student",
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err != nil {
		panic("err init client:" + err.Error())
	}

	resp, err := cli.Register(context.Background(), &demo.Student{
		Id:   1,
		Name: "Zhong",
		Sex:  "male",
		College: &demo.College{
			Name:    "NJU",
			Address: "JS",
		},
		Email: []string{
			"211180179@smail.nju.edu.cn",
		},
	})
	if err != nil {
		panic("register failed")
	}

	klog.Infof("resp: %v", resp)
}
