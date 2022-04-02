package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"column:username; unique;" form:"username" binding:"required"`                      //default string length = varchar(255)
	Password     string `gorm:"column:password;type:varbinary(100);" json:"-" form:"password" binding:"required"` //sha512로 저장예정
	Name         string `gorm:"column:name" form:"name"`                                                          //form - ShouldBind에서 쓰기 위한 것
	ProfileImage string `gorm:"column:profile_image"`
	Room         []Room
	RoomUser     []RoomUser
}

//테이블 이름 지정
func (u *User) TableName() string {
	return "User"
}

//아이디 중복체크
func (u *User) CheckIsUnique() {

}
