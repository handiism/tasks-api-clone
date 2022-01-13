package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint
	ListId    uint           `gorm:"not null"`
	Name      string         `gorm:"type:varchar(50);not null"`
	Detail    sql.NullString `gorm:"type:varchar(100)"`
	DueDate   sql.NullTime
	IsDone    bool      `gorm:"default:false;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:current_timestamp;not null"`
	UpdatedAt sql.NullTime
	Subtasks  []Subtask
}

func (Task) TableName() string {
	return "task"
}
