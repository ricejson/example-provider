package service

import (
	"github.com/ricejson/example-common/model"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (u *UserServiceImpl) GetUser() (model.User, error) {
	return model.User{Name: "ricejson"}, nil
}
