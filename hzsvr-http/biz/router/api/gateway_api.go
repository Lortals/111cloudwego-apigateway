// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	api "github.com/Lortals/111cloudwego-apigateway/hzsvr-http/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/add", append(_addnumbersMw(), api.AddNumbers)...)
	root.POST("/add-student-info", append(_registerMw(), api.Register)...)
	root.POST("/divide", append(_dividenumbersMw(), api.DivideNumbers)...)
	root.POST("/multiply", append(_multiplynumbersMw(), api.MultiplyNumbers)...)
	root.GET("/query", append(_queryMw(), api.Query)...)
}
