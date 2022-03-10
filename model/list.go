package model

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uint      `json:"id" faker:"-"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id" faker:"-"`
	Title     string    `gorm:"size:255;not null" json:"title" faker:"word"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" faker:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" faker:"-"`
	Tasks     []Task    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tasks" faker:"-"`
}

func (List) TableName() string {
	return "list"
}
