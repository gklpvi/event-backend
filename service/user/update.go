package userServices

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
)

func Update(input *inputs.UpdateUserInput) error {
	var user models.User

	if result := models.DB.First(&user, input.Id); result.Error != nil {
		return result.Error
	}
	user.Verified = input.Verified

	models.DB.Save(&user)
	return nil
}
