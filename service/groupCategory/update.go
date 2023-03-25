package projectServices

import (
	"event-backend/model"
	"event-backend/model/inputs"
)

func Update(input *inputs.UpdateStatusInput) error {
	var status model.Status

	if result := model.DB.First(&status, input.Id); result.Error != nil {
		return result.Error
	}
	status.Name = input.Name
	status.Type = input.Type

	model.DB.Save(&status)
	return nil
}
