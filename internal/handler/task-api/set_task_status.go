package task_api

import (
	pb "ToDoList/generated/api/task"
	"context"
)

func (a TaskApi) SetTaskStatus(ctx context.Context, req *pb.SetTaskStatusRequest) (*pb.SetTaskStatusResponse, error) {

	if err := a.TaskService.SetTaskStatus(ctx, req.TaskId, req.TaskStatus); err != nil {
		return nil, err
	}
	return &pb.SetTaskStatusResponse{}, nil
}
