package userServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateUserInput) error {
	var user model.User

	if result := model.DB.First(&user, input.Id); result.Error != nil {
		return result.Error
	}

	model.DB.Save(&user)
	return nil
}
