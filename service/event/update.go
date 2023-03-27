package eventServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateGroupInput) error {
	var event model.Group
	if result := model.DB.First(&event, input.Id); result.Error != nil {
		return result.Error
	}

	event.GroupCategoryID = input.GroupCategoryID
	event.MaxMember = input.MaxMember

	model.DB.Save(&event)
	return nil
}
