package models

import (
	"database/sql"

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

type RoomUsers struct {
	gorm.Model
	RoomName sql.NullString `gorm:"null"`
	RoomType sql.NullInt16  `gorm:"null; default:1"`
	Owner    User           `gorm:"foreignKey:UserID"`
	UserID   uint
	Notice   bool `gorm:"default:false"`
}

func (ru RoomUsers) TableName() string {
	return "RoomUsers"
}
