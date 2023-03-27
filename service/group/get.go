package groupServices

import "event-backend/model"

func GetById(id uint) (*model.Group, error) {
	var group model.Group
	if result := model.DB.Preload("GroupMembers").First(&group, id); result.Error != nil {
		return &model.Group{}, result.Error
	}
	return &group, nil
}

func GetAll() ([]model.Group, error) {
	var group []model.Group
	if result := model.DB.Preload("GroupMembers").Find(&group); result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func GetGroupMembers(groupID uint) ([]model.Profile, error) {
	var group model.Group
	if result := model.DB.Preload("GroupMembers").Preload("GroupMembers.Profile").First(&group, groupID); result.Error != nil {
		return nil, result.Error
	}
	var profiles []model.Profile
	for _, groupMember := range group.GroupMembers {
		profiles = append(profiles, groupMember.Profile)
	}
	return profiles, nil
}

func GetByPlayerId(playerId uint, eventId uint) (*model.Group, error) {
	var group model.Group
	if result := model.DB.Preload("GroupMembers").Where("event_id=?", eventId).First(&group); result.Error != nil {
		return &model.Group{}, result.Error
	}
	for _, groupMember := range group.GroupMembers {
		if groupMember.ProfileID == playerId {
			return &group, nil
		}
	}
	return &model.Group{}, nil
}
