package groupServices

import (
	"event-backend/model"
)

func IsGroupFull(groupID uint) (bool) {
	var group model.Group
	if result := model.DB.Preload("GroupMembers").First(&group, groupID); result.Error != nil {
		return false
	}

	if len(group.GroupMembers) >= group.MaxMember {
		return true
	}

	return false
}

func GetByLevel(level int, eventID uint64) (*model.Group, error){
	var group model.Group
	if result := model.DB.Where("level=? AND event_id=?", level, eventID).First(&group); result.Error != nil {
		return &model.Group{}, result.Error
	}

	return &group, nil
}