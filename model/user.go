package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id" faker:"-"`
	Name      string    `gorm:"size:255;not null" json:"name" faker:"name"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email" faker:"email"`
	Password  string    `gorm:"size:255;not null" json:"-" faker:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" faker:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" faker:"-"`
	Token     string    `gorm:"-" json:"token,omitempty" faker:"-"`
	Lists     []List    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"list" faker:"-"`
}

func (User) TableName() string {
	return "user"
}
