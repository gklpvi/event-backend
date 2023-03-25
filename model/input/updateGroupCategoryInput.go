package input

type UpdateGroupCategoryInput struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	MaxLevel int    `json:"max_level"`
}
