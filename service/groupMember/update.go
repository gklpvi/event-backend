package groupMemberServices

import (
	"event-backend/model"
	"event-backend/model/input"
)

func Update(input *input.UpdateGroupMemberInput) error {
	var groupMember model.GroupMember

	if result := model.DB.First(&groupMember, input.Id); result.Error != nil {
		return result.Error
	}

	groupMember.GroupID = input.GroupID
	groupMember.ProfileID = input.ProfileID

	model.DB.Save(&groupMember)
	return nil
}
