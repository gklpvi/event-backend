package model

import "github.com/jinzhu/gorm"

type GroupCategory struct {
	gorm.Model
	Name 		string `gorm:"size:255;not null;" json:"name"`
	MaxLevel 	int    `gorm:"not null;" json:"max_level"`
}
