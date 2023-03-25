package input

type UpdateUserInput struct{
	Id            uint   `json:"id"`
	Mail          string `json:"mail"`
	Password      string `json:"password"`
}