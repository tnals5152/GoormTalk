package models

import (
	"gorm.io/gorm"
)

type RoomTypeDomainStruct struct {
	OneToOne int
	Group    int
}

var RoomTypeDomain RoomTypeDomainStruct = RoomTypeDomainStruct{
	OneToOne: 1,
	Group:    2,
}

type RoomUser struct {
	gorm.Model
	Room   Room `gorm:"foreignKey:RoomID"`
	RoomID uint
	User   User `gorm:"foreignKey:UserID"`
	UserID uint
	Notice bool `gorm:"default:false"`
}

func (ru RoomUser) TableName() string {
	return "RoomUser"
}
