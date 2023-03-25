package projectServices

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
)

func Update(input *inputs.UpdateStatusInput) error {
	var status models.Status

	if result := models.DB.First(&status, input.Id); result.Error != nil {
		return result.Error
	}
	status.Name = input.Name
	status.Type = input.Type

	models.DB.Save(&status)
	return nil
}
