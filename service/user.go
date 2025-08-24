package service

import (
	"fmt"
	"github.com/ricejson/example-common/model"
)

type UserServiceImpl struct {
}

func (u *UserServiceImpl) getUser(user model.User) (model.User, error) {
	// 打印名字
	fmt.Println(user.GetName())
	return user, nil
}
