package service

import (
	"fmt"
	"github.com/ricejson/example-common/model"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) GetUser(user model.User) (model.User, error) {
	fmt.Println(user.GetName())
	return user, nil
}
