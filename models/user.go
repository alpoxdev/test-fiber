package models

import (
	"time"

	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(255);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nickname      string `gorm:"not null;unique" json:"nickname"`
	Email     string `gorm:"not null;unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}

// BeforeCreate는 GORM 훅으로, 레코드가 생성되기 전에 호출됩니다.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = cuid2.Generate()
	return	
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
