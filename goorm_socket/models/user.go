package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"column:username"` //default string length = varchar(255)
	Password     string `gorm:"column:password"` //sha512로 저장예정
	Name         string `gorm:"column:name"`
	ProfileImage string `gorm:"column:profile_image"`
	Room         []Room
}

//테이블 이름 지정
func (u User) TableName() string {
	return "User"
}