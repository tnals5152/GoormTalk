package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomName string `gorm:"not null; default:room"` //room_name
	RoomType uint
	// Owner    User `gorm:"foreignKey:UserID"`
	UserID uint
}

func (r Room) TableName() string {
	return "Room"
}
