package groupCategoryServices

import "event-backend/model"

func GetById(id uint) (*model.GroupCategory, error) {

	var groupCategory model.GroupCategory
	if result := model.DB.First(&groupCategory, id); result.Error != nil {
		return &model.GroupCategory{}, result.Error
	}
	return &groupCategory, nil
}
func GetAll() ([]model.GroupCategory, error) {

	var groupCategory []model.GroupCategory
	if result := model.DB.Find(&groupCategory); result.Error != nil {
		return nil, result.Error
	}
	return groupCategory, nil
}
