package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:20;not null"`
	Pass  string `gorm:"size:20;not null"`
	Phone int64  `gorm:"not null"`
}

var user User

func IsExist(name string) bool {
	GetUserByName(name)
	return user.ID != 0
}

func GetUserByPhone(phone int64) *User {
	db.Where("phone=?", phone).First(&user)
	return &user
}

func GetUserByName(name string) *User {
	db.Where("name=?", name).First(&user)
	return &user
}

func RegUser(name, pwd string, phone int64) error {
	user = User{
		Name:  name,
		Pass:  pwd,
		Phone: phone,
	}
	return db.Create(&user).Error
}
