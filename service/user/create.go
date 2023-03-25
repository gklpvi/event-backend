package userServices

import "gradpanel-backend/models"

func Create(user *models.User) (*models.User, error) {

	if result := models.DB.Create(&user); result.Error != nil {
		return &models.User{}, result.Error
	}

	return user, nil
}
