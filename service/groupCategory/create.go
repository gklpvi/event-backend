package projectServices

import "event-backend/model"

func Create(status *model.Status) (*model.Status, error) {

	if result := model.DB.Create(&status); result.Error != nil {
		return &model.Status{}, result.Error
	}

	return status, nil
}
