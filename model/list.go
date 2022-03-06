package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uint         `gorm:"type:SERIAL;->" json:"id" faker:"-"`
	UserID    uuid.UUID    `gorm:"type:uuid;not null" json:"user_id" faker:"-"`
	Title     string       `gorm:"size:255;not null" json:"title" faker:"word"`
	CreatedAt time.Time    `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"default:null" json:"updated_at" faker:"-"`
	Tasks     []Task       `gorm:"foreignKey:ListID" json:"tasks" faker:"-"`
}

func (List) TableName() string {
	return "list"
}
