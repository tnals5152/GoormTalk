package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomName string //room_name
	RoomType uint
}
