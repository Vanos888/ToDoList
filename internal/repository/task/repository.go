package task

import (
	"ToDoList/internal/domain"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const taskTableName = "task"

type TaskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (r *TaskRepo) CreateTask(ctx context.Context, task domain.Task) error {
	q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(taskTableName).
		Columns("user_id", "task_name", "priority", "task_desc").
		Values(
			task.UserID,
			task.TaskName,
			task.Priority,
			task.TaskDesc,
		)

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *TaskRepo) SetTaskStatus(ctx context.Context, taskId string, status string) error {
	q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(taskTableName).
		Set("task_status", status).
		Where(sq.Eq{"task_id": taskId})

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepo) GetTaskByFilter(ctx context.Context, filter domain.TaskFilter) ([]domain.Task, error) {
	q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("task_id", "user_id", "task_name", "task_desc", "task_status", "priority", "created_at").
		From(taskTableName).
		Where(sq.Eq{"user_id": filter.UserID})

	if filter.Priority != 0 {
		q = q.Where(sq.Eq{"priority": filter.Priority})
	}
	if filter.TaskStatus != "" {
		q = q.Where(sq.Eq{"task_status": filter.TaskStatus})
	}
	if filter.OrderField != domain.OrderFieldUnknown {
		q = q.OrderBy(filter.OrderField.String() + " " + filter.OrderBy.String())
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		task := domain.Task{}
		err := rows.Scan(&task.TaskID, &task.UserID, &task.TaskName, &task.TaskDesc, &task.TaskStatus, &task.Priority, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}
