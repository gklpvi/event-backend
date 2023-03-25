package projectServices

import "event-backend/model"

func GetById(id uint) (*model.Status, error) {

	var status model.Status
	if result := model.DB.First(&status, id); result.Error != nil {
		return &model.Status{}, result.Error
	}
	return &status, nil
}
func GetAll() ([]model.Status, error) {

	var status []model.Status
	if result := model.DB.Find(&status); result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}
