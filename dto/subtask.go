package dto

type CreateSubtask struct {
	TaskId uint   `json:"task_id" binding:"required" form:"task_id"`
	Name   string `json:"name" form:"name" binding:"required" validate:"min:1,max:50"`
}

type UpdateSubtask struct {
	ID     uint   `json:"id" form:"id" binding:"required"`
	TaskId uint   `json:"task_id" binding:"required" form:"task_id"`
	Name   string `json:"name" form:"name" binding:"required" validate:"min:1,max:50"`
	IsDone bool   `json:"is_done" form:"is_done"`
}
