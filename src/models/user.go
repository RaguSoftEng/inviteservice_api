package models

import "gorm.io/gorm"

// All users are considered admin users and to reduce the complexity, properties are limited to username and password only.
type User struct {
	gorm.Model
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `gorm:"size:100;not null;" json:"password"`
}
