package task

import (
	"ToDoList/internal/domain"
	"context"
)

type ITaskARepo interface {
	CreateTask(ctx context.Context, task domain.Task) error
	SetTaskStatus(ctx context.Context, taskId string, status string) error
	GetTaskByFilter(ctx context.Context, filter domain.TaskFilter) ([]domain.Task, error)
}

type Service struct {
	taskRepo ITaskARepo
}

func NewTaskService(taskRepo ITaskARepo) *Service {
	return &Service{
		taskRepo: taskRepo,
	}
}

func (s *Service) CreateTask(ctx context.Context, task domain.Task) error {
	return s.taskRepo.CreateTask(ctx, task)
}
func (s *Service) SetTaskStatus(ctx context.Context, taskId string, status string) error {
	return s.taskRepo.SetTaskStatus(ctx, taskId, status)
}

func (s *Service) GetTaskByFilter(ctx context.Context, filter domain.TaskFilter) ([]domain.Task, error) {
	return s.taskRepo.GetTaskByFilter(ctx, filter)
}
