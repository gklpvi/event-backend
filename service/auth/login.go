package auth

import (
	"event-backend/model"
)

func LoginService(mail string) (*model.User, error) {

	user := model.User{}
	if result := model.DB.Model(model.User{}).Where("mail=?", mail).Take(&user); result.Error != nil {
		return &model.User{}, result.Error
	}

	return &user, nil
}
