package dto

type CreateUser struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1,max:50"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email,max:50"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8,max:50,alphanum"`
}

type UpdateUser struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required" validate:"required,min:1,max:50"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email,max:50"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8,max:50,alphanum"`
}
