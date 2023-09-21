package _usecase

import (
	"context"
	"errors"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/domain"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/internal/helper"
)

type UserUsacaseImpl struct {
	userRepo domain.UserRepository
}

func NewUserUsacaseImpl(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsacaseImpl{
		userRepo: userRepo,
	}
}

func (u *UserUsacaseImpl) Create(ctx context.Context, req *domain.RequestCreateUser) (*domain.ResponseUser, error) {
	passwordHash, err := helper.Hashing(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	}
	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	resp := &domain.ResponseUser{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}
func (u *UserUsacaseImpl) Login(ctx context.Context, req *domain.RequestLogin) (*domain.ResponseLogin, error) {
	user, err := u.userRepo.GetByEmailOrID(ctx, req.EmailOrUsername, req.EmailOrUsername)
	if err != nil {
		return nil, err
	}

	ok := helper.Compare(req.Password, user.Password)
	if !ok {
		return nil, errors.New("invalid password")
	}

	jwtModel := &helper.Jwt{
		UserID: user.ID,
		Key:    "random key from backend",
		Exp:    5 * time.Hour,
	}
	token, err := helper.GenerateJwtHS256(jwtModel)
	if err != nil {
		return nil, err
	}

	resp := &domain.ResponseLogin{
		Token: token,
		ResponseUser: domain.ResponseUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	return resp, nil
}
