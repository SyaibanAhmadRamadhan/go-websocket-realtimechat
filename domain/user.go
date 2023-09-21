package domain

import (
	"context"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

type RequestCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestLogin struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}

type ResponseUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseLogin struct {
	Token        string `json:"token"`
	ResponseUser `json:"user"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (int64, error)
	GetByEmailOrID(ctx context.Context, username, email string) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, req *RequestCreateUser) (*ResponseUser, error)
	Login(ctx context.Context, req *RequestLogin) (*ResponseLogin, error)
}
