package model

import "github.com/jinzhu/gorm"

type GroupCategory struct {
	gorm.Model
	Name 		string `gorm:"size:255;not null;" json:"name"`
	MinLevel 	int    `gorm:"not null;" json:"minLevel"`
	MaxLevel 	int    `gorm:"not null;" json:"maxLevel"`
}
