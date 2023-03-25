package userServices

import "gradpanel-backend/models"


func GetById(id uint) (*models.User, error) {

	var user models.User
	if result := models.DB.First(&user, id); result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}

func GetAll() ([]models.User, error) {

	var user []models.User
	if result := models.DB.Find(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetByMail(mail string) (*models.User, error) {
	var user models.User
	if result := models.DB.Where("mail=?", mail).Take(&user); result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}
