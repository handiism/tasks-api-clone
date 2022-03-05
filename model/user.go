package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `gorm:"type:uuid;primaryKey" json:"id" faker:"-"`
	Name      string       `gorm:"type:varchar(50);not null" json:"name" faker:"name"`
	Email     string       `gorm:"uniqueIndex;type:varchar(50);not null" json:"email" faker:"email"`
	Password  string       `gorm:"type:varchar(50);not null" json:"-" faker:"password"`
	CreatedAt time.Time    `gorm:"not null" json:"created_at" faker:"-"`
	UpdatedAt sql.NullTime `gorm:"default:null" json:"updated_at" faker:"-"`
	Token     string       `gorm:"-" json:"token,omitempty" faker:"-"`
	Lists     []List       `faker:"-"`
}

func (User) TableName() string {
	return "user"
}
