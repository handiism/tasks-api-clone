package model

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uint      `json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Tasks     []Task    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"tasks,omitempty"`
}

func (List) TableName() string {
	return "list"
}
