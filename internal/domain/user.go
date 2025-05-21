package domain

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	TgName       string    `json:"tg_name" db:"tg_name"`
	PasswordHash string    `json:"pass_hash" db:"pass_hash"`
}
type Login struct {
	TgName   string `json:"tg_name" db:"tg_name"`
	Password string `json:"pass_hash" db:"pass_hash"`
}

func (u User) IsEmpty() bool {
	if u.ID == uuid.Nil {
		return true
	}
	return false
}
