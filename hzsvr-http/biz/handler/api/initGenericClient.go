package api

import (
	"log"
	"os"

	kitexClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// 泛化调用函数实现
func initGenericClient() genericclient.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	//支持热更新功能的IDL
	path := "../IDL-manage/gateway_api.thrift"
	cont, err := os.ReadFile(path)
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
