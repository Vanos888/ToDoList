package auth_api

import (
	pb "ToDoList/generated/api/auth"
	"ToDoList/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (a *AuthApi) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		ID:           uuid.New(),
		TgName:       req.User.TgName,
		PasswordHash: req.User.Password,
	}

	err := a.AuthService.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{}, nil

}
