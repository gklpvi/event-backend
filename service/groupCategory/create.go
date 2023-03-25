package groupCategoryServices

import "event-backend/model"

func Create(groupCategory *model.GroupCategory) (*model.GroupCategory, error) {

	if result := model.DB.Create(&groupCategory); result.Error != nil {
		return &model.GroupCategory{}, result.Error
	}

	return groupCategory, nil
}
