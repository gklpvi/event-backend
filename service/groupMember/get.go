package projectServices

import "gradpanel-backend/models"

func GetById(id uint) (*models.ProjectStudent, error) {

	var projectStudent models.ProjectStudent
	if result := models.DB.First(&projectStudent, id); result.Error != nil {
		return &models.ProjectStudent{}, result.Error
	}
	return &projectStudent, nil
}
func GetAll() ([]models.ProjectStudent, error) {

	var projectStudent []models.ProjectStudent
	if result := models.DB.Find(&projectStudent); result.Error != nil {
		return nil, result.Error
	}
	return projectStudent, nil
}
