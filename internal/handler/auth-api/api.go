package auth_api

import (
	pb "ToDoList/generated/api/auth"
	"ToDoList/internal/domain"
	"context"
	"github.com/google/uuid"
)

type IService interface {
	Register(ctx context.Context, user domain.User) error
	ConfirmCode(ctx context.Context, tgName string, code int) (uuid.UUID, error)
	Login(ctx context.Context, user domain.Login) (string, error)
}

type AuthApi struct {
	pb.UnimplementedAuthApiServer
	AuthService IService
}

func NewAuthApi(AuthService IService) *AuthApi {
	return &AuthApi{
		AuthService: AuthService,
	}
}
