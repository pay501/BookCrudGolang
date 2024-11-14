package model

import "spy/repository"

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	SignUp(*repository.User) error
	Login(*LoginReq) (string, error)
}
