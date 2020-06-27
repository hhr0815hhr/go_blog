package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Account   string `gorm:"size:20;not null"`
	Name      string `gorm:"size:20;not null"`
	Pass      string `gorm:"size:80;not null"`
	Email     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var user User

func IsExist(name string) bool {
	GetUserByAccount(name)
	return user.ID != 0
}

func GetUserByAccount(account string) *User {
	db.Where("account=?", account).First(&user)
	return &user
}

func GetUserById(userId uint) *User {
	db.First(&user, userId)
	return &user
}

func RegUser(name, pwd, email string) (error, *User) {
	user = User{
		Account: name,
		Name:    name,
		Pass:    pwd,
		Email:   email,
	}
	return db.Create(&user).Error, &user
}
