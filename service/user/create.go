package userServices

import "event-backend/model"

func Create(user *model.User) (*model.User, error) {

	if result := model.DB.Create(&user); result.Error != nil {
		return &model.User{}, result.Error
	}

	return user, nil
}
