package models

import (
	"crypto/sha512"
	"goorm_socket/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"column:username; unique;" form:"username" binding:"required"`                      //default string length = varchar(255)
	Password     string `gorm:"column:password;type:varbinary(100);" json:"-" form:"password" binding:"required"` //sha512로 저장예정
	Name         string `gorm:"column:name" form:"name"`                                                          //form - ShouldBind에서 쓰기 위한 것
	ProfileImage string `gorm:"column:profile_image"`
	Room         []Room
	RoomUser     []RoomUser
}

var PasswordLength int = 8

//테이블 이름 지정
func (u *User) TableName() string {
	return "User"
}

//아이디 중복체크
func (user *User) CheckIsUnique() bool {
	var count int64
	config.GetDB.Model(user).Where("username = ?", user.Username).Count(&count)

	if count > 0 {
		return false
	}
	return true
}

func (user *User) LoginCheck() (*gorm.DB, []User) {
	var users []User
	result := config.GetDB.Model(user).Where(user).Find(&users)
	return result, users
}

func (user *User) ChangePassword() {
	password := sha512.Sum512([]byte(user.Password))
	user.Password = string(password[:])
}

func (user *User) CreateUser() *gorm.DB {
	return config.SetDB.Model(&user).Create(&user)
}
