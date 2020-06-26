package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:20;not null"`
	Pass  string `gorm:"size:20;not null"`
	Phone int64  `gorm:"not null"`
}
