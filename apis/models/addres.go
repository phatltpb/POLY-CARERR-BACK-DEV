package models

import "time"

type City struct {
	Id        int       `gorm:"primary_key;column:id" json:"id"`
	Parent    string    `gorm:"column:parent" json:"parent"`
	Name      string    `gorm:"column:name" json:"name"`
	Code      string    `gorm:"column:code" json:"code"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
