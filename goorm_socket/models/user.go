package models

type User struct {
	// gorm.Model
	User_ID      uint64 `gorm:"column:user_id; primary_key"`
	Username     string `gorm:"column:username"`
	Name         string `gorm:"column:name"`
	ProfileImage string `gorm:"column:profile_image"`
}
