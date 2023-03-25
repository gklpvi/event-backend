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
