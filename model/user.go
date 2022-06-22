package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"           json:"id"`
	Name      string    `gorm:"size:255;not null"                             json:"name"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"                 json:"email"`
	Password  string    `gorm:"size:255;not null"                             json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                                json:"updated_at"`
	Token     string    `gorm:"-"                                             json:"token,omitempty"`
	Lists     []List    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"list,omitempty"`
}

func (User) TableName() string {
	return "user"
}
