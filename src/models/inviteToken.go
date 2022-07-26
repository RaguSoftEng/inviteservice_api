package models

import (
	"gorm.io/gorm"
)

type InviteToken struct {
	gorm.Model
	CreaterId uint   `gorm:"not null" json:"-"`
	CreatedBy User   `gorm:"foreignKey:CreaterId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"createdBy"`
	Token     string `json:"token"`
	IsEnabled bool   `json:"isEnabled"`
}
