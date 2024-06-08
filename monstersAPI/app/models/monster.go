package models

import (
	"time"

	"gorm.io/gorm"
)

type Monster struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null;size:255" json:"name" validate:"required"`
	Attack    uint           `gorm:"not null" json:"attack" validate:"required"`
	Defense   uint           `gorm:"not null" json:"defense" validate:"required"`
	Hp        uint           `gorm:"not null" json:"hp" validate:"gte=0"`
	Speed     uint           `gorm:"not null" json:"speed" validate:"required"`
	ImageURL  string         `gorm:"not null" json:"imageUrl"`
	Battles   []Battle       `gorm:"foreignKey:ID" json:"battles"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
