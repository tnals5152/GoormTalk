package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	UrlPath   string  `gorm:"type:text"`
	Message   Message `gorm:"foreignKey:MessageID"`
	MessageID uint
}

func (l Link) TableName() string {
	return "Link"
}
