package groupMemberServices

import (
	"event-backend/model"
)

func HasJoinedEvent(playerID uint, eventID uint64) (bool) {
	var groupMember model.GroupMember
	if result := model.DB.Where("player_id=? AND event_id=?", playerID, eventID).First(&groupMember); result.Error != nil {
		return false
	}

	return true
}