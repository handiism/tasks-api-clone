package model

import (
	"database/sql"
	"time"
)

type Subtask struct {
	ID        uint
	TaskId    uint         `gorm:"not null"`
	Name      string       `gorm:"type:varchar(50);not null"`
	IsDone    bool         `gorm:"default:false;not null"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`
}

func (Subtask) TableName() string {
	return "subtask"
}
