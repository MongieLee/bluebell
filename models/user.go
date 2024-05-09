package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint64         `json:"id"`
	UserId    int64          `json:"user_id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Gender    int            `json:"gender" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Token     string         `gorm:"-"`
}
