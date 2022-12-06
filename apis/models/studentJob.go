package models

import "time"

type StudentCV struct {
	Id        int        `gorm:"column:id" json:"id"`
	StudentId int        `gorm:"column:student_id" json:"student_id"`
	Title     string     `gorm:"title" json:"title"`
	Link      string     `gorm:"link" json:"link"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	DeleteAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	Students *Students `gorm:"foreignKey:StudentId" json:"student,omitempty"`
}

type ApplyJob struct {
	Id        int        `gorm:"column:id" json:"id"`
	PostId    int        `gorm:"column:post_id" json:"job_id"`
	CvId      int        `gorm:"column:cv_id" json:"cv_id"`
	Status    int        `gorm:"column:status" json:"status"`
	Letter    string     `gorm:"column:letter" json:"letter"`
	DateApply time.Time  `gorm:"column:date_apply" json:"date_apply"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`

	Job       *Job       `gorm:"foreignKey:PostId" json:"job,omitempty"`
	StudentCV *StudentCV `gorm:"foreignKey:CvId" json:"cv,omitempty"`
}

type CVFilter struct {
	JobId    int
	DateFrom *time.Time
	DateTo   *time.Time
	Status   int
	Page     int
	Limit    int
}
