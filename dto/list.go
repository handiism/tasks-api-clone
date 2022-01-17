package dto

type CreateList struct {
	UserID uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
	Title  string `json:"title" form:"tittle" binding:"required" validate:"min:1,max:50"`
}

type UpdateList struct {
	ID     uint   `json:"id" form:"id" binding:"required"`
	UserID uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
	Title  string `json:"title" form:"title" binding:"required" validate:"min:1,max:50"`
}
