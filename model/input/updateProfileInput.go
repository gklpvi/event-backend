package inputs

type UpdateProfileInput struct{
	Id            uint   `json:"id"`
	Username	  string `json:"username"`
	Level		 int	`json:"level"`
}