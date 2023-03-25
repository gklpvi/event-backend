package model

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Username string `json:"username"`
	Level    int    `json:"level"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
}
