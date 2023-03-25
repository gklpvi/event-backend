package projectServices

import "gradpanel-backend/models"

func Delete(id uint) error {
	if result := models.DB.Delete(&models.ProjectStudent{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
