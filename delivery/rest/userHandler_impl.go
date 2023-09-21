package rest

import (
	"encoding/json"
	"net/http"

	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/delivery/rest/helper"
	"github.com/SyaibanAhmadRamadhan/go-websocket-realtimechat/domain"
)

type UserHandlerImpl struct {
	userUsecase domain.UserUsecase
}

func NewUserHandlerImpl(userUsecase domain.UserUsecase) *UserHandlerImpl {
	return &UserHandlerImpl{
		userUsecase: userUsecase,
	}
}

func (u *UserHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateUser)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	resp, err := u.userUsecase.Create(r.Context(), req)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	helper.SuccessEncode(w, resp, "successfully created user")
}

func (u *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestLogin)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	resp, err := u.userUsecase.Login(r.Context(), req)
	if err != nil {
		helper.ErrorEncode(w, err)
		return
	}

	helper.SuccessEncode(w, resp, "successfully login")
}
