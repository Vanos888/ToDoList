package auth_api

import (
	pb "ToDoList/generated/api/auth"
	"ToDoList/internal/domain"
	"context"
	"crypto/sha256"
	"encoding/base64"
)

func (a *AuthApi) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	reqPass := sha256.New()
	reqPass.Write([]byte(req.User.Password))
	reqPassHash := base64.URLEncoding.EncodeToString(reqPass.Sum(nil))

	user := domain.Login{
		TgName:   req.User.TgName,
		Password: reqPassHash,
	}

	jwt, err := a.AuthService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Jwt: jwt,
	}, nil
}
