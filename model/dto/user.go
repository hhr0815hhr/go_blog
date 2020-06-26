package dto

import "blog/model"

type UserDto struct {
	Name  string
	Phone int64
}

func User(user *model.User) *UserDto {
	return &UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
