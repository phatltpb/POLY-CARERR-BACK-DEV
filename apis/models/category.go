package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int            `gorm:"primary_key;column:id" json:"id"`
	ParentId  int            `gorm:"column:parent_id" json:"parent_id"`
	Name      string         `gorm:"column:name" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"colum:deleted_at" json:"deleted_at"`
}
