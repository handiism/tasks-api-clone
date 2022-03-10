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
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at" faker:"-"`
	Subtasks  []Subtask      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"subtask" faker:"-"`
}

func (Task) TableName() string {
	return "task"
}
