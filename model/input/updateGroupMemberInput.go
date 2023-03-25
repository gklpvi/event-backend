package input

type UpdateGroupMemberInput struct {
	Id        uint `json:"id"`
	GroupID   uint `json:"group_id"`
	ProfileID uint `json:"profile_id"`
}
