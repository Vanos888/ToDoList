package task_api

import (
	pb "ToDoList/generated/api/task"
	"ToDoList/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (a *TaskApi) GetTaskByFilter(ctx context.Context, req *pb.GetTaskByFilterRequest) (*pb.GetTaskByFilterResponse, error) {
	userID, err := uuid.Parse(req.Filter.UserId)
	if err != nil {
		return nil, err
	}

	filter := domain.TaskFilter{
		UserID:     userID,
		TaskStatus: req.Filter.TaskStatus,
		Priority:   req.Filter.Priority,
		OrderField: domain.OrderField(req.Filter.OrderField),
		OrderBy:    domain.OrderBy(req.Filter.OrderBy),
	}

	tasks, err := a.TaskService.GetTaskByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetTaskByFilterResponse{}

	for _, task := range tasks {
		resp.Tasks = append(resp.Tasks, &pb.Task{
			TaskId:     task.TaskID.String(),
			UserId:     task.UserID.String(),
			TaskName:   task.TaskName,
			TaskDesk:   task.TaskDesc,
			TaskStatus: task.TaskStatus,
			Priority:   task.Priority,
			CreatedAt:  task.CreatedAt.String(),
		})
	}

	return resp, nil

}
