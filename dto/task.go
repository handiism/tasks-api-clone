package dto

import "database/sql"

type CreateTask struct {
	ListId  uint           `json:"list_id" binding:"required" form:"list_id"`
	Name    string         `json:"name" form:"name" binding:"required" validate:"min:1,max:50"`
	Detail  sql.NullString `json:"detail" form:"name"`
	DueDate sql.NullTime   `json:"due_date" form:"due_date"`
}

type UpdateTask struct {
	ID      uint           `json:"id" form:"id" binding:"required"`
	ListId  uint           `json:"list_id" binding:"required" form:"list_id"`
	Name    string         `json:"name" form:"name" binding:"required" validate:"min:1,max:50"`
	Detail  sql.NullString `json:"detail" form:"name"`
	DueDate sql.NullTime   `json:"due_date" form:"due_date"`
	IsDone  bool           `json:"is_done" form:"is_done"`
}
