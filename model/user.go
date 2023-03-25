package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Mail     string `gorm:"size:255;not null;" json:"mail" validate:"required,email"`
	Password string `gorm:"size:255;not null;" json:"password" validate:"required"`
}
