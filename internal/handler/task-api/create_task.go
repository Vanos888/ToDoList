package task_api

import (
	pb "ToDoList/generated/api/task"
	"ToDoList/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (a TaskApi) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {

	userID, err := uuid.Parse(req.Task.UserId)
	if err != nil {
		return nil, err
	}
	task := domain.Task{
		UserID:   userID,
		TaskName: req.Task.TaskName,
		Priority: req.Task.Priority,
		TaskDesc: req.Task.TaskDesc,
	}
	err = a.TaskService.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskResponse{}, nil
}
