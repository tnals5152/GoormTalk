package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Path      string `gorm:"type:text; not null"`
	Width     float64
	Height    float64
	FileType  string  `grom:"type:varchar(30)"` //==size:30
	Message   Message `grom:"foreignKey:MessageID"`
	MessageID uint
}

func (f File) TableName() string {
	return "File"
}
