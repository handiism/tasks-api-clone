package model

import (
	"time"
)

type Subtask struct {
	ID        uint
	TaskID    uint      `gorm:"not null" json:"task_id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	IsDone    bool      `gorm:"default:false;not null" json:"is_done"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Subtask) TableName() string {
	return "subtask"
}
