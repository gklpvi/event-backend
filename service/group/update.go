package projectServices

import (
	"event-backend/model"
)

func Update(input *inputs.UpdateProjectInput) error {
	var project model.Project
	if result := model.DB.First(&project, input.Id); result.Error != nil {
		return result.Error
	}

	project.Title = input.Title
	project.Major = input.Major
	project.Capacity = input.Capacity
	project.Status = input.Status
	project.ProfileID = input.InstructorID

	model.DB.Save(&project)
	return nil
}
