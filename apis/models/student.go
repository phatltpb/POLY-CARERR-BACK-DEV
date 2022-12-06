package models

import (
	"time"
)

type Students struct {
	Id         int        `gorm:"primary_key;auto_increment" json:"id"`
	Email      string     `gorm:"column:email" json:"email"`
	FullName   string     `gorm:"column:full_name" json:"full_name"`
	Password   string     `gorm:"column:password" json:"password"`
	Avatar     string     `gorm:"column:avatar" json:"avatar"`
	LocationId *int       `gorm:"column:location" json:"location_id"`
	Location   *City      `gorm:"foreignKey:LocationId" json:"location"`
	Phone      string     `gorm:"column:phone" json:"phone"`
	Address    *string    `gorm:"column:address" json:"address"`
	Birthday   *time.Time `gorm:"column:birthday" json:"birthday"`
	Gender     int        `gorm:"column:gender" json:"gender"`
	IsNoti     bool       `gorm:"column:is_noti" json:"is_noti"`
	IsActive   bool       `gorm:"column:is_active" json:"is_active"`
	Status     int        `gorm:"column:status" json:"status"`
	Role       int        `gorm:"column:role" json:"role"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`

	StudentProfile   *StudentProfile     `gorm:"foreignKey:StudentId" json:"profile"`
	StudentEducation []*StudentEducation `gorm:"foreignKey:StudentId" json:"education"`
}

type StudentProfile struct {
	Id           int    `gorm:"column:id"  json:"-"`
	StudentId    int    `gorm:"column:student_id" json:"student_id"`
	PositionWish string `gorm:"column:position_wish" json:"position_wish"`
	LevelWish    string `gorm:"column:level_wish" json:"level_wish"`
	LevelCurrent string `gorm:"column:level_current" json:"level_current"`
	Experience   string `gorm:"column:experience" json:"experience"`
	SalaryWish   int64  `gorm:"column:salary_wish" json:"salary_wish"`
	ProvinceId   int    `gorm:"column:province_id" json:"province_id"`
	CategoryWish int    `gorm:"column:category_id" json:"category_wish"`
	JobType      string `gorm:"column:job_type" json:"job_type_wish"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Province *City     `gorm:"foreignKey:ProvinceId" json:"province"`
	Category *Category `gorm:"foreignKey:CategoryWish" json:"category"`
}

type StudentEducation struct {
	Id          int    `json:"id"`
	StudentId   int    `gorm:"column:student_id" json:"student_id"`
	Degree      string `gorm:"column:degree" json:"degree"`
	Rank        string `gorm:"column:rank" json:"rank"`
	Information string `gorm:"column:information" json:"information"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type StudentWatch struct {
	Id        int   `gorm:"id"`
	StudentId int   `gorm:"column:student_id"`
	JobId     int   `gorm:"column:job_id"`
	Count     int64 `gorm:"column:count"`
}
