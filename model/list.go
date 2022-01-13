package model

import (
	"database/sql"
	"time"
)

type List struct {
	ID        uint
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:current_timestamp;not null"`
	UpdatedAt sql.NullTime
	Tasks     []Task
}

func (List) TableName() string {
	return "list"
}
