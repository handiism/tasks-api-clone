package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid()" json:"id" faker:"-"`
	Name      string       `gorm:"size:255;not null" json:"name" faker:"name"`
	Email     string       `gorm:"uniqueIndex;size:255;not null" json:"email" faker:"email"`
	Password  string       `gorm:"size:255;not null" json:"-" faker:"password"`
	CreatedAt time.Time    `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"default:null" json:"updated_at" faker:"-"`
	Token     string       `gorm:"-" json:"token,omitempty" faker:"-"`
	Lists     []List       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"list" faker:"-"`
}

func (User) TableName() string {
	return "user"
}
