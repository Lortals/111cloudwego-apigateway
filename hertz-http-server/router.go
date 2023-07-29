// Code generated by hertz generator.

package main

import (
	handler "github.com/Lortals/111cloudwego-apigateway/hertz-http-server/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
