package user

import (
	"ToDoList/internal/domain"
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const userTableName = "users"

type UserRepo struct {
	db *pgxpool.Pool
}

var columns = []string{
	"id",
	"tg_name",
	"pass_hash",
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error) {
	q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(userTableName).
		Columns(columns...).
		Values(user.ID, user.TgName, user.PasswordHash)

	query, args, err := q.ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *UserRepo) GetUserByLogin(ctx context.Context, tgName string) (domain.User, error) {
	//reqPass := sha256.New()
	//reqPass.Write([]byte(user.Password))
	//reqPassHash := reqPass.Sum(nil)

	sql, args, err := sq.
		Select("id", "tg_name", "pass_hash").
		From(userTableName).
		Where(sq.Eq{
			"tg_name": tgName,
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return domain.User{}, err
	}

	row, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, nil
		}

		return domain.User{}, err
	}

	for row.Next() {
		res := domain.User{}
		if err := row.Scan(&res.ID, &res.TgName, &res.PasswordHash); err != nil {
			return domain.User{}, err
		}

		return res, nil
	}

	return domain.User{}, nil
}
