package projectServices

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
)

func Update(input *inputs.UpdateProjectInput) error {
	var project models.Project
	if result := models.DB.First(&project, input.Id); result.Error != nil {
		return result.Error
	}

	project.Title = input.Title
	project.Major = input.Major
	project.Capacity = input.Capacity
	project.Status = input.Status
	project.ProfileID = input.InstructorID

	models.DB.Save(&project)
	return nil
}
