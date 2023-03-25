package groupServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateGroupInput) error {
	var group model.Group
	if result := model.DB.First(&group, input.Id); result.Error != nil {
		return result.Error
	}

	group.GroupCategoryID = input.GroupCategoryID
	group.MaxMember = input.MaxMember

	model.DB.Save(&group)
	return nil
}
