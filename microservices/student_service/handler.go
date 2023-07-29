package main

import (
	"context"
	api "github.com/ararchch/api-gateway/microservices/student_service/kitex_gen/student/api"
)

// StudentManagementImpl implements the last service interface defined in the IDL.
type StudentManagementImpl struct{}

// RegisterStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) RegisterStudent(ctx context.Context, req *api.RegisterStudentRequest) (resp *api.RegisterStudentResponse, err error) {
	return &api.RegisterStudentResponse{
		Status: "success",
		Id: "1",
	}, nil
}

// GetStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) GetStudent(ctx context.Context, req *api.GetStudentRequest) (resp *api.GetStudentResponse, err error) {
	return &api.GetStudentResponse{
		Name: "Xiao Hong",
		Age: "18",
		Email: "xiaohong@tsinghua.edu.cn",
		Address: "Tsinghua University, 30 Shuangqing Road, Haidian District, Beijing 100084, China.",
	}, nil
}
