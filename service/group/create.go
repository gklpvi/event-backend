package projectServices

import "gradpanel-backend/models"

func Create(project *models.Project) (*models.Project, error) {
	if result := models.DB.Create(&project); result.Error != nil {
		return &models.Project{}, result.Error
	}

	return project, nil
}
