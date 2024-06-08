package models

import (
	"time"

	"gorm.io/gorm"
)

type Battle struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	MonsterAID uint           `gorm:"column:monster_a" json:"-"`
	MonsterA   Monster        `gorm:"not null;foreignKey:MonsterAID" json:"monsterA" validate:"required"`
	MonsterBID uint           `gorm:"column:monster_b" json:"-"`
	MonsterB   Monster        `gorm:"not null;foreignKey:MonsterBID" json:"monsterB" validate:"required"`
	WinnerID   uint           `gorm:"column:winner" json:"-"`
	Winner     Monster        `gorm:"not null;foreignKey:WinnerID" json:"winner" validate:"required"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
