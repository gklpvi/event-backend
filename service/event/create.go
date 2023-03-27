package eventServices

import "event-backend/model"

func Create(event *model.Event) (*model.Event, error) {
	if result := model.DB.Create(&event); result.Error != nil {
		return &model.Event{}, result.Error
	}

	return event, nil
}
