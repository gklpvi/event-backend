package groupServices

import "event-backend/model"

func Create(group *model.Group) (*model.Group, error) {
	if result := model.DB.Create(&group); result.Error != nil {
		return &model.Group{}, result.Error
	}

	return group, nil
}
