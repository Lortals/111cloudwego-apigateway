// Code generated by Kitex v0.6.1. DO NOT EDIT.

package multiplicationmanagement

import (
	management "github.com/Lortals/111cloudwego-apigateway/rpc-services/multiplication-service/kitex_gen/multiplication/management"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler management.MultiplicationManagement, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
