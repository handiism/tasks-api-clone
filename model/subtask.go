package model

import (
	"database/sql"
	"time"
)

type Subtask struct {
	ID        uint
	TaskId    uint      `gorm:"not null"`
	Name      string    `gorm:"type:varchar(50);not null"`
	IsDone    bool      `gorm:"default:false;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:current_timestamp;not null"`
	UpdatedAt sql.NullTime
}

func (Subtask) TableName() string {
	return "subtask"
}
