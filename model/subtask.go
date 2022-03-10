package model

import (
	"time"
)

type Subtask struct {
	ID        uint      `faker:"-"`
	TaskID    uint      `gorm:"not null" json:"task_id" faker:"ref"`
	Name      string    `gorm:"size:255;not null" json:"name" faker:"word"`
	IsDone    bool      `gorm:"default:false;not null" json:"is_done" faker:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" faker:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" faker:"-"`
}

func (Subtask) TableName() string {
	return "subtask"
}
