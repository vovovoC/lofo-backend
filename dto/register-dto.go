package dto

//RegisterDTO is used when client POST from /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" validate:"min:6"`
}
