package projectServices

import "gradpanel-backend/models"

func GetById(id uint) (*models.Project, error) {
	var project models.Project
	if result := models.DB.Preload("Profile").First(&project, id); result.Error != nil {
		return &models.Project{}, result.Error
	}
	return &project, nil
}

func GetAll() ([]models.Project, error) {
	var project []models.Project
	if result := models.DB.Preload("Profile").Find(&project); result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func GetDetails(id uint) (*models.Project, error) {
	var project models.Project
	if err := models.DB.Preload("Files").Preload("Files.Profile").Preload("Comments").Preload("Comments.Profile").Preload("Application").Preload("Application.Profile").Preload("Source").Preload("Source.Profile").Preload("Link").Preload("Link.Profile").Preload("Profile").Preload("Requirement").Preload("Requirement.Profile").Preload("ProjectStudent").Preload("ProjectStudent.Profile").First(&project, id).Error; err != nil {
		return &models.Project{}, err
	}

	return &project, nil
}
