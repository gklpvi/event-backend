package projectServices

import "event-backend/model"

func Create(projectStudent *model.ProjectStudent) (*model.ProjectStudent, error) {

	if result := model.DB.Create(&projectStudent); result.Error != nil {
		return &model.ProjectStudent{}, result.Error
	}

	return projectStudent, nil
}
