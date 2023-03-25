package groupCategoryServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateGroupCategoryInput) error {
	var groupCategory model.GroupCategory

	if result := model.DB.First(&groupCategory, input.Id); result.Error != nil {
		return result.Error
	}
	groupCategory.Name = input.Name
	groupCategory.MaxLevel = input.MaxLevel

	model.DB.Save(&groupCategory)
	return nil
}
