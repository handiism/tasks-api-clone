package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint           `json:"id"`
	ListID    uint           `json:"list_id"           gorm:"not null"`
	Name      string         `json:"name"              gorm:"size:255;not null"`
	Detail    sql.NullString `json:"detail"            gorm:"size:255"`
	DueDate   sql.NullTime   `json:"due_date"          gorm:"default:null"`
	IsDone    bool           `json:"is_done"           gorm:"default:false;not null"`
	CreatedAt time.Time      `json:"created_at"        gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at"        gorm:"autoUpdateTime"`
	Subtasks  []Subtask      `json:"subtask,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Task) TableName() string {
	return "task"
}
