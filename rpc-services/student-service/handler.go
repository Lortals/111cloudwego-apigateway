package main

import (
	"context"
	"errors"
	"sync"

	management "github.com/Lortals/111cloudwego-apigateway/rpc-services/student-service/kitex_gen/student/management"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	mu sync.RWMutex // Define a mutex to protect id2Student
}

// map 缓存
var id2Student = map[int]management.Student{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *management.Student) (resp *management.RegisterResp, err error) {
	// TODO: Your code here...
	s.mu.Lock()         // Lock the mutex before modifying the map
	defer s.mu.Unlock() // Unlock the mutex after the function exits

	// rpc register
	id := int(student.Id)
	_, found := id2Student[id]
	if found {
		resp = &management.RegisterResp{
			Success: false,
			Message: "This ID already exists.",
		}
		return
	}

	id2Student[id] = *student
	resp = &management.RegisterResp{
		Success: true,
		Message: "Information added successfully.",
	}

	// fmt.Println(id2Student)
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *management.QueryReq) (resp *management.Student, err error) {
	// TODO: Your code here...
	s.mu.Lock()         // Lock the mutex before modifying the map
	defer s.mu.Unlock() // Unlock the mutex after the function exits

	student, found := id2Student[int(req.Id)]
	if found {
		resp = &student
	} else {
		err = errors.New("not found")
	}

	// fmt.Println(resp)

	return
}
