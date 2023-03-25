package model

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Groups       []Group       `gorm:"foreignkey:EventID" json:"groups"`
	GroupMembers []GroupMember `gorm:"foreignkey:EventID" json:"groupMembers"`
}
