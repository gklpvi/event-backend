package groupMemberServices

import (
	"event-backend/model"
)

func HasJoinedEvent(playerID uint, eventID uint64) bool {
	var groupMember model.GroupMember
	var group []model.Group
	var groupIds []uint

	// get groups based on eventId
	if result := model.DB.Where("event_id=?", eventID).Find(&group); result.Error != nil {
		return false
	}
	// get list of groupIds
	for _, g := range group {
		groupIds = append(groupIds, g.ID)
	}
	// get groupMember based on playerId and groupIds
	if result := model.DB.Where("profile_id=? AND group_id IN (?)", playerID, groupIds).First(&groupMember); result.Error != nil {
		return false
	}

	return true
}
