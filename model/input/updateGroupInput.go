package input

type UpdateGroupInput struct {
	Id        uint `json:"id"`
	GroupCategoryID  uint `json:"group_category_id"`
	MaxMember int `json:"max_member"`
}
