package userServices

import "gradpanel-backend/models"

func Delete(id uint) error {
	if result := models.DB.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
