package projectServices

import "gradpanel-backend/models"

func GetById(id uint) (*models.Status, error) {

	var status models.Status
	if result := models.DB.First(&status, id); result.Error != nil {
		return &models.Status{}, result.Error
	}
	return &status, nil
}
func GetAll() ([]models.Status, error) {

	var status []models.Status
	if result := models.DB.Find(&status); result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}
