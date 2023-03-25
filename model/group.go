package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	EventID 	   uint          `gorm:"not null;" json:"eventId"`
	Event          Event         `gorm:"foreignkey:eventID" json:"event"`
	GroupCategoryID uint          `gorm:"not null;" json:"groupCategoryId"`
	GroupCategory   GroupCategory `gorm:"foreignkey:groupCategoryID" json:"groupCategory"`
	MaxMember       int           `gorm:"not null;" json:"maxMember"`
	GroupMembers    []GroupMember `gorm:"foreignkey:groupID" json:"groupMembers"`
}
