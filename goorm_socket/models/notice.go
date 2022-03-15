package models

import (
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Room      Room `gorm:"foreignKey:RoomID"`
	RoomID    uint
	Owner     User `gorm:"foreignKey:UserID"`
	UserID    uint
	Message   Message `gorm:"foreignKey:MessageID"`
	MessageID uint
}

func (n Notice) TableName() string {
	return "Notice"
}
