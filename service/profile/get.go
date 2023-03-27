package profileServices

import "event-backend/model"

// GetById returns a single profile from the database by ID
func GetById(id uint) (*model.Profile, error) {
	var profile model.Profile
	// First return the first record that match the condition
	if result := model.DB.Preload("User").First(&profile, id); result.Error != nil {
		// return empty profile on error
		return &model.Profile{}, result.Error
	}
	return &profile, nil
}

// GetByUserId returns a single profile from the database by user ID
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
