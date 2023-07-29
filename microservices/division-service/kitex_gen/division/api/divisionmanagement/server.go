// Code generated by Kitex v0.6.1. DO NOT EDIT.
package divisionmanagement

import (
	api "github.com/ararchch/api-gateway/microservices/division-service/kitex_gen/division/api"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.DivisionManagement, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
