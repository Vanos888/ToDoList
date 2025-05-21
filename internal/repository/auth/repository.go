package auth

import (
	"ToDoList/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type AuthRepo struct {
	db *redis.Client
}

func NewAuthRepo(db *redis.Client) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) SaveTempUser(ctx context.Context, user domain.User, code int) error {
	userDto := AuthUserDto{
		PassHash: user.PasswordHash,
		Code:     code,
	}

	fmt.Println(user.TgName)

	return r.db.Set(ctx, user.TgName, userDto, 5*time.Minute).Err()
}

func (r *AuthRepo) GetTempUser(ctx context.Context, tgName string) (domain.User, int, error) {
	res, err := r.db.Get(ctx, tgName).Result()
	if err != nil {
		return domain.User{}, 0, err
	}

	user := AuthUserDto{}
	if err := json.Unmarshal([]byte(res), &user); err != nil {
		return domain.User{}, 0, err
	}

	return domain.User{TgName: tgName, PasswordHash: user.PassHash}, user.Code, nil
}
