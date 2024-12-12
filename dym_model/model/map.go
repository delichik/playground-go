package model

import (
	"time"
)

type Map struct {
	ID        int       `gorm:"primaryKey;default:" json:"id"`
	Version   string    `gorm:"size:64;not null;default:" json:"version"`
	Content   string    `gorm:"not null;default:" json:"content"`
	Status    int       `gorm:"not null;default:" json:"status"`
	Remark    string    `gorm:"size:1024;not null;default:" json:"remark"`
	CreatedAt time.Time `gorm:"not null;default:" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:" json:"updated_at"`
}
