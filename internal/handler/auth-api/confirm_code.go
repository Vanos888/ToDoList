package auth_api

import (
	pb "ToDoList/generated/api/auth"
	"context"
)

func (a *AuthApi) ConfirmCode(ctx context.Context, req *pb.ConfirmCodeRequest) (*pb.ConfirmCodeResponse, error) {
	id, err := a.AuthService.ConfirmCode(ctx, req.TgName, int(req.Code))
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmCodeResponse{
		UserId: id.String(),
	}, nil

}
