package profileServices

import "event-backend/model"

func GetById(id uint) (*model.Profile, error) {

	var profile model.Profile
	if result := model.DB.Preload("User").First(&profile, id); result.Error != nil {
		return &model.Profile{}, result.Error
	}
	return &profile, nil
}

func GetByUserId(user_id uint) (*model.Profile, error) {
	var profile model.Profile
	if result := model.DB.Model(model.Profile{}).Where("user_id=?", user_id).Take(&profile); result.Error != nil {
		return &model.Profile{}, result.Error
	}
	return &profile, nil
}

func GetAll() ([]model.Profile, error) {
	var profile []model.Profile
	if result := model.DB.Preload("User").Find(&profile); result.Error != nil {
		return nil, result.Error
	}
	return profile, nil
}
