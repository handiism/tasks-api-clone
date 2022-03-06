package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint           `json:"id" faker:"-"`
	ListID    uint           `gorm:"not null" json:"list_id" faker:"ref"`
	Name      string         `gorm:"size:255;not null" json:"name" faker:"word"`
	Detail    sql.NullString `gorm:"size:255" json:"detail" faker:"-"`
	DueDate   sql.NullTime   `gorm:"default:null" json:"due_date" faker:"-"`
	IsDone    bool           `gorm:"default:false;not null" json:"is_done" faker:"-"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime   `gorm:"default:null" json:"updated_at" faker:"-"`
	Subtasks  []Subtask      `gorm:"foreignKey:TaskID;references:ID" json:"subtask" faker:"-"`
}

func (Task) TableName() string {
	return "task"
}
