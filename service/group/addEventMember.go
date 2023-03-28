package groupServices

import (
	"event-backend/model"
)

func IsGroupFull(groupID uint) (bool) {
	var group model.Group
	if result := model.DB.Preload("GroupMembers").First(&group, groupID); result.Error != nil {
		return false
	}
	// return true if group has reached its max member capacity
	return len(group.GroupMembers) >= group.MaxMember
}

func GetByLevel(level int, eventID uint64) (*model.Group, error) {
	var group model.Group

	// levels for the group categories are hardcoded since the document says so but in the future it can be dynamic since the model and database already support it
	groupCategoryId := func(level int) uint {
		if level >= 0 && level < 20 {
			return 1
		} else if level >= 20 && level < 50 {
			return 2
		} else {
			return 3
		}
	}(level)

	// get last group by group category id and event id
	if result := model.DB.Where("group_category_id=? AND event_id=?", groupCategoryId, eventID).Last(&group); result.Error != nil {
		return &model.Group{}, result.Error
	}

	return &group, nil
}