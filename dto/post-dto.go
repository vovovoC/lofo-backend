package dto

//PostUpdateDTO is a model that client use when update a post
type PostUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	Image       string `json:"image" form:"image"`
	Place       string `json:"place" form:"place" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required"`
	Status      string `json:"status" form:"status" binding:"required"`
	Datetime    string `json:"datetime" form:"datetime" binding:"required"`
}

//PostCreateDTO is a model that client use when create a new post
type PostCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	Image       string `json:"image" form:"image"`
	Place       string `json:"place" form:"place" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required"`
	Status      string `json:"status" form:"status" binding:"required"`
	Datetime    string `json:"datetime" form:"datetime" binding:"required"`
}
