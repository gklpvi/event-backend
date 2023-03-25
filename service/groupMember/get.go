package groupMemberServices

import "event-backend/model"

func GetById(id uint) (*model.GroupMember, error) {

	var groupMember model.GroupMember
	if result := model.DB.First(&groupMember, id); result.Error != nil {
		return &model.GroupMember{}, result.Error
	}
	return &groupMember, nil
}
func GetAll() ([]model.GroupMember, error) {

	var groupMember []model.GroupMember
	if result := model.DB.Find(&groupMember); result.Error != nil {
		return nil, result.Error
	}
	return groupMember, nil
}
