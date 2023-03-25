package userServices

import "event-backend/model"

func GetById(id uint) (*model.User, error) {

	var user model.User
	if result := model.DB.First(&user, id); result.Error != nil {
		return &model.User{}, result.Error
	}
	return &user, nil
}

func GetAll() ([]model.User, error) {

	var user []model.User
	if result := model.DB.Find(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetByMail(mail string) (*model.User, error) {
	var user model.User
	if result := model.DB.Where("mail=?", mail).Take(&user); result.Error != nil {
		return &model.User{}, result.Error
	}
	return &user, nil
}
