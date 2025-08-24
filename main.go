package main

import (
	"fmt"
	"github.com/ricejson/example-common/model"
	"github.com/ricejson/example-common/service"
)

func main() {
	// todo 获取到UserService的实现类对象
	var userService service.UserService
	user := model.User{}
	user.SetName("rice")
	newUser, err := userService.GetUser(user)
	if err == nil {
		fmt.Println(newUser.GetName())
	} else {
		fmt.Println(err)
	}
}
