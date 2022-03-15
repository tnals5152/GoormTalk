package models

import (
	"gorm.io/gorm"
)

type MessageTypeDomainStruct struct {
	Message int
	File    int
	Link    int
}

var MessageTypeDomain MessageTypeDomainStruct = MessageTypeDomainStruct{
	Message: 1,
	File:    2,
	Link:    3,
}

type Message struct {
	gorm.Model
	Room        Room `gorm:"foreignKey:RoomID"`
	RoomID      uint
	User        User `gorm:"foreignKey:UserID"`
	UserID      uint
	Content     string `gorm:"type:text"`
	MessageType uint   `gorm:"not null; default:1"`
}

func (m Message) TableName() string {
	return "Message"
}
