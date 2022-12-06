package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Id          int    `gorm:"primary_key;column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	CategoryId  int    `gorm:"column:category_id" json:"category_id,omitempty"`
	UserId      int    `gorm:"column:user_id" json:"user_id"`
	Gender      string `gorm:"column:gender" json:"gender,omitempty"`
	LocationId  *int   `gorm:"column:location" json:"province_id,omitempty"`
	Address     string `gorm:"column:address" json:"address,omitempty"`
	JobType     string `gorm:"column:job_type" json:"job_type,omitempty"`
	Count       int    `gorm:"count" json:"count,omitempty"`
	Experience  string `gorm:"experience" json:"experience,omitempty"`
	Position    string `gorm:"position" json:"position,omitempty"`
	Salary      int64  `gorm:"salary" json:"salary,omitempty"`
	Level       string `gorm:"level" json:"level,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Require     string `gorm:"column:require" json:"require,omitempty"`
	Benefit     string `gorm:"column:benefit" json:"benefit,omitempty"`
	IsHidden    bool   `gorm:"column:is_hidden" json:"is_hidden"`
	Status      int    `gorm:"column:status" json:"status"`

	Location *City     `gorm:"foreignKey:LocationId" json:"province,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryId" json:"category,omitempty"`
	User     *User     `gorm:"foreignKey:UserId" json:"author,omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdateAt  *time.Time     `gorm:"column:updated_at" json:"update_at"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type JobFilter struct {
	CompanyId  int
	CategoryId int
	Search     string
	MinSalary  int
	MaxSalary  int
	Level      string
	Experience string
	Position   string

	Pagination struct {
		Page  int
		Limit int
	}
}
