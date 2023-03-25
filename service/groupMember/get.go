package projectServices

import "event-backend/model"

func GetById(id uint) (*model.ProjectStudent, error) {

	var projectStudent model.ProjectStudent
	if result := model.DB.First(&projectStudent, id); result.Error != nil {
		return &model.ProjectStudent{}, result.Error
	}
	return &projectStudent, nil
}
func GetAll() ([]model.ProjectStudent, error) {

	var projectStudent []model.ProjectStudent
	if result := model.DB.Find(&projectStudent); result.Error != nil {
		return nil, result.Error
	}
	return projectStudent, nil
}
