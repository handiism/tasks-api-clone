package model

import (
	"database/sql"
	"time"
)

type Subtask struct {
	ID        uint         `faker:"-"`
	TaskID    uint         `gorm:"not null" json:"task_id" faker:"ref"`
	Name      string       `gorm:"size:255;not null" json:"name" faker:"word"`
	IsDone    bool         `gorm:"default:false;not null" json:"is_done" faker:"-"`
	CreatedAt time.Time    `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"default:null" json:"updated_at" faker:"-"`
}

func (Subtask) TableName() string {
	return "subtask"
}
