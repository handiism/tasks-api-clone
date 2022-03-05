package model

import (
	"database/sql"
	"time"
)

type List struct {
	ID        uint         `json:"id"`
	UserID    string       `gorm:"not null" json:"user_id" faker:"-"`
	Title     string       `gorm:"type:varchar(50);not null" json:"title" faker:"word"`
	CreatedAt time.Time    `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"default:null" json:"updated_at" faker:"-"`
	Tasks     []Task       `faker:"-"`
}

func (List) TableName() string {
	return "list"
}
