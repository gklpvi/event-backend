package eventServices

import "event-backend/model"

func GetById(id uint) (*model.Event, error) {
	var event model.Event
	if result := model.DB.Preload("Groups").Preload("GroupMembers").First(&event, id); result.Error != nil {
		return &model.Event{}, result.Error
	}
	return &event, nil
}

func GetAll() ([]model.Event, error) {
	var event []model.Event
	if result := model.DB.Preload("Groups").Preload("GroupMembers").Find(&event); result.Error != nil {
		return nil, result.Error
	}
	return event, nil
}
