package projectServices

import "gradpanel-backend/models"

func Create(projectStudent *models.ProjectStudent) (*models.ProjectStudent, error) {

	if result := models.DB.Create(&projectStudent); result.Error != nil {
		return &models.ProjectStudent{}, result.Error
	}

	return projectStudent, nil
}
