package model

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uint      `json:"id"`
	UserID    uuid.UUID `json:"user_id"         gorm:"type:uuid;not null"`
	Title     string    `json:"title"           gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at"      gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"      gorm:"autoUpdateTime"`
	Tasks     []Task    `json:"tasks,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (List) TableName() string {
	return "list"
}
