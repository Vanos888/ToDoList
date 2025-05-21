package task_api

import (
	pb "ToDoList/generated/api/task"
	"ToDoList/internal/domain"
	"context"
)

type IService interface {
	CreateTask(ctx context.Context, task domain.Task) error
	SetTaskStatus(ctx context.Context, taskId string, status string) error
	GetTaskByFilter(ctx context.Context, filter domain.TaskFilter) ([]domain.Task, error)
}

type TaskApi struct {
	pb.UnimplementedTaskApiServer
	TaskService IService
}

func NewTaskApi(taskService IService) *TaskApi {
	return &TaskApi{
		TaskService: taskService,
	}
}
