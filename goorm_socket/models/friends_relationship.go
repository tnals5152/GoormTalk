package models

import "gorm.io/gorm"

type FriendsRelationship struct {
	gorm.Model
	User     User `gorm:"foreignKey:UserID"`
	UserID   uint
	Friend   User `gorm:"foreignkey:FriendID"`
	FriendID uint
}

func (f FriendsRelationship) TableName() string {
	return "FriendsRelationship"
}
