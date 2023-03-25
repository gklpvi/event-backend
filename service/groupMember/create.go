package groupMemberServices

import "event-backend/model"

func Create(groupMember *model.GroupMember) (*model.GroupMember, error) {

	if result := model.DB.Create(&groupMember); result.Error != nil {
		return &model.GroupMember{}, result.Error
	}

	return groupMember, nil
}
