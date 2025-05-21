package domain

import (
	"github.com/google/uuid"
	"time"
)

type OrderField int

const (
	OrderFieldUnknown OrderField = iota
	OrderFieldCreatedAt
	OrderFieldPriority
)

var orderFieldMap = map[OrderField]string{
	OrderFieldUnknown:   "unknown",
	OrderFieldCreatedAt: "created_at",
	OrderFieldPriority:  "priority",
}

func (o OrderField) String() string {
	if v, ok := orderFieldMap[o]; ok {
		return v
	}

	return "unknown"
}

type OrderBy int

const (
	OrderByAsc OrderBy = iota
	OrderByDesc
)

var orderByMap = map[OrderBy]string{
	OrderByAsc:  "ASC",
	OrderByDesc: "DESC",
}

func (o OrderBy) String() string {
	if v, ok := orderByMap[o]; ok {
		return v
	}

	return "unknown"
}

type Task struct {
	UserID     uuid.UUID `json:"user_id"`
	TaskID     uuid.UUID `json:"task_id"`
	TaskName   string    `json:"task_name"`
	TaskDesc   string    `json:"task_desc"`
	TaskStatus string    `json:"task_status"`
	Priority   int64     `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TaskFilter struct {
	UserID     uuid.UUID  `json:"user_id"`
	TaskStatus string     `json:"task_status"`
	Priority   int64      `json:"priority"`
	OrderField OrderField `json:"order_field"`
	OrderBy    OrderBy    `json:"order_by"`
}
