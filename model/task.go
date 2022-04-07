package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint           `json:"id"`
	ListID    uint           `gorm:"not null" json:"list_id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Detail    sql.NullString `gorm:"size:255" json:"detail"`
	DueDate   sql.NullTime   `gorm:"default:null" json:"due_date"`
	IsDone    bool           `gorm:"default:false;not null" json:"is_done"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Subtasks  []Subtask      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"subtask,omitempty"`
}

func (Task) TableName() string {
	return "task"
}
