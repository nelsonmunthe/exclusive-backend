package entity

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	UserId       string    `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Password     string    `json:"password"`
	Phone_Number string    `json:"phone_number"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}
