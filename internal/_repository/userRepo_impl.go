package _repository

import (
	"context"
	"database/sql"

	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (int64, error) {
	query := `INSERT INTO users(username, email, password) VALUES($1, $2, $3) returning id`

	conn, err := u.db.Conn(ctx)
	if err != nil {
		return 0, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	var id int
	err = stmt.QueryRowContext(ctx, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func (u *UserRepositoryImpl) GetByEmailOrID(ctx context.Context, username, email string) (*domain.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1 OR email = $2`

	conn, err := u.db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = stmt.QueryRowContext(ctx, username, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
