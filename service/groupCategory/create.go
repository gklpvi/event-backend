package projectServices

import "gradpanel-backend/models"

func Create(status *models.Status) (*models.Status, error) {

	if result := models.DB.Create(&status); result.Error != nil {
		return &models.Status{}, result.Error
	}

	return status, nil
}
