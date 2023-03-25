package model

import "github.com/jinzhu/gorm"

type GroupMember struct {
	gorm.Model
	GroupID   uint   `gorm:"not null;" json:"group_id"`
	Group     Group  `gorm:"foreignkey:groupID" json:"group"`
	ProfileID uint   `gorm:"not null;" json:"profile_id"`
	Profile   Profile `gorm:"foreignkey:profileID" json:"profile"`
}
