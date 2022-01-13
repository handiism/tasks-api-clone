package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	Name      string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"uniqueIndex;type:varchar(50);not null"`
	Password  string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:current_timestamp;not null"`
	UpdatedAt sql.NullTime
	Lists     []List
}

func (User) TableName() string {
	return "user"
}
