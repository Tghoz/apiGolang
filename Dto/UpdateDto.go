package dto

type UpdateTagsRequest struct {

  
	UserName string `validate:"required,max=200,min=1" json:"user_name"`
	Email    string `json:"email" validate:"required,email"`


}