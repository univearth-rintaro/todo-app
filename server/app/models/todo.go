package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      string         `gorm:"type:varchar(50);not null" json:"status"`
	Done        bool           `gorm:"default:false" json:"done"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
