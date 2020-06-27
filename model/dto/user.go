package dto

import "blog/model"

type UserDto struct {
	Account string
	Name    string
}

func User(user *model.User) *UserDto {
	return &UserDto{
		Account: user.Account,
		Name:    user.Name,
	}
}
