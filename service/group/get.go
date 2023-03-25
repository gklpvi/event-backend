package projectServices

import "event-backend/model"

func GetById(id uint) (*model.Project, error) {
	var project model.Project
	if result := model.DB.Preload("Profile").First(&project, id); result.Error != nil {
		return &model.Project{}, result.Error
	}
	return &project, nil
}

func GetAll() ([]model.Project, error) {
	var project []model.Project
	if result := model.DB.Preload("Profile").Find(&project); result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func GetDetails(id uint) (*model.Project, error) {
	var project model.Project
	if err := model.DB.Preload("Files").Preload("Files.Profile").Preload("Comments").Preload("Comments.Profile").Preload("Application").Preload("Application.Profile").Preload("Source").Preload("Source.Profile").Preload("Link").Preload("Link.Profile").Preload("Profile").Preload("Requirement").Preload("Requirement.Profile").Preload("ProjectStudent").Preload("ProjectStudent.Profile").First(&project, id).Error; err != nil {
		return &model.Project{}, err
	}

	return &project, nil
}
