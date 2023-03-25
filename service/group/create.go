package projectServices

import "event-backend/model"

func Create(project *model.Project) (*model.Project, error) {
	if result := model.DB.Create(&project); result.Error != nil {
		return &model.Project{}, result.Error
	}

	return project, nil
}
