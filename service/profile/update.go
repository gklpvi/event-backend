package profileServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateProfileInput) error {
	var profile model.Profile

	if result := model.DB.First(&profile, input.Id); result.Error != nil {
		return result.Error
	}

	profile.Username = input.Username
	profile.Level = input.Level

	model.DB.Save(&profile)
	return nil
}
