package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupCategoryID uint          `gorm:"not null;" json:"group_category_id"`
	GroupCategory   GroupCategory `gorm:"foreignkey:groupCategoryID" json:"group_category"`
	MaxMember       int           `gorm:"not null;" json:"max_member"`
	GroupMembers    []GroupMember `gorm:"foreignkey:groupID" json:"group_members"`
}
