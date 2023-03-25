package profileServices

import "event-backend/model"

func Create(profile *model.Profile) (*model.Profile, error) {

	if result := model.DB.Create(&profile); result.Error != nil {
		return &model.Profile{}, result.Error
	}

	return profile, nil
}
