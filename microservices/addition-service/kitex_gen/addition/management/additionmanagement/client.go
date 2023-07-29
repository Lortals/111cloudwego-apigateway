// Code generated by Kitex v0.6.1. DO NOT EDIT.

package additionmanagement

import (
	"context"
	management "github.com/ararchch/api-gateway/microservices/addition-service/kitex_gen/addition/management"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddNumbers(ctx context.Context, req *management.AdditionRequest, callOptions ...callopt.Option) (r *management.AdditionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAdditionManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAdditionManagementClient struct {
	*kClient
}

func (p *kAdditionManagementClient) AddNumbers(ctx context.Context, req *management.AdditionRequest, callOptions ...callopt.Option) (r *management.AdditionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNumbers(ctx, req)
}
