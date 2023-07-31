package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func GenerateClient(serviceName string, opts ...kitexClient.Option) (genericclient.Client, error) {

	// inital declarations
	var err error

	// initating loadbalancer
	lb := loadbalance.NewWeightedBalancer()

	// initating etcs resolver (for service discovery)
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	// importing idl for reference(generic call)
	p, err := generic.NewThriftFileProvider("../IDL-manage/gateway_api.thrift")
	if err != nil {
		panic(err)
	}

	// convert to thrift generic form
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	var options []kitexClient.Option
	options = append(options,
		kitexClient.WithResolver(r),
		kitexClient.WithLoadBalancer(lb),
	)
	options = append(options, opts...)

	// create new generic client
	client, err := genericclient.NewClient(
		serviceName,
		g,
		options...,
	)
	if err != nil {
		panic(err)
	}

	return client, nil
}

func jsonStringify(item any) (string, error) {
	// convert to request struct to JSON format (so it can be converted to json string)
	jsonForm, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	return string(jsonForm), nil
}

func MakeRpcRequest(ctx context.Context, kitexClient genericclient.Client, methodName string, request interface{}, response interface{}) error {
	stringedReq, err := jsonStringify(request)
	if err != nil {
		panic(err)
	}

	// making generic call to addNumbers method of client
	respRpc, err := kitexClient.GenericCall(ctx, methodName, stringedReq)
	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(respRpc.(string)), response)

	return nil
}

// 泛化调用
func initGenericClient() genericclient.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	//基于内存解析 IDL，支持热更新
	path := "../IDL-manage/gateway_api.thrift"
	cont, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(cont[:])

	includes := map[string]string{
		path: content,
	}

	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	}

	// dynamic update
	err = p.UpdateIDL(content, includes)
	if err != nil {
		panic("UpdateIDL failed")
	}

	// 构造HTTP类型的泛化调用
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	cli, err := genericclient.NewClient("student", g,
		kitexClient.WithResolver(r),
		kitexClient.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err != nil {
		panic(err)
	}

	return cli
}
