package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"column:username"`
	Password     string `gorm:"column:password"`
	Name         string `gorm:"column:name"`
	ProfileImage string `gorm:"column:profile_image"`
}

//테이블 이름 지정
func (u User) TableName() string {
	return "User"
}
