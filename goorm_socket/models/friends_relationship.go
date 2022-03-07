package models

import "gorm.io/gorm"

type FriendsRelationship struct {
	gorm.Model
	User   User `gorm:"foreignKey:UserID"`
	UserID uint
	// Friend User //`gorm:"foreignkey:ID"`
}

func (f FriendsRelationship) TableName() string {
	return "FriendsRelationship"
}
