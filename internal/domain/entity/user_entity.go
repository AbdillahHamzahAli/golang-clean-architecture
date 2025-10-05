package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username string    `gorm:"type:varchar(255)" json:"user_name"`
	Email    string    `gorm:"type:varchar(255)" json:"email"`
	Password string    `gorm:"type:varchar(255)" json:"password"`

	Posts     []Post    `gorm:"foreignkey:UserID" json:"posts"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}
