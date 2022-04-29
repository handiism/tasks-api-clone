package request

import "database/sql"

type SaveUser struct {
	Name     string `binding:"required"                json:"name"`
	Email    string `binding:"required,email"          json:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password"`
}

type SaveList struct {
	Title string `binding:"required,min=1" json:"title"`
}

type SaveTask struct {
	Name    string         `binding:"required"         json:"name"`
	Detail  sql.NullString `binding:"min=0"            json:"detail,omitempty"`
	DueDate sql.NullTime   `binding:"datetime"         json:"due_date,omitempty"`
	IsDone  bool           `binding:"required,boolean" json:"is_done"`
}

type SaveSubtask struct {
	Name   string `binding:"required"         json:"name"`
	IsDone bool   `binding:"required,boolean" json:"is_done"`
}
