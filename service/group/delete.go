package projectServices

import "event-backend/model"

func Delete(id uint) error {
	if result := model.DB.Delete(&model.Project{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
