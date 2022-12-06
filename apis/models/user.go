package models

import "time"

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	EmailNoti string    `json:"email_noti"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Avatar    string    `json:"avatar"`
	Is_noti   bool      `json:"is_noti"`
	Role      int       `json:"role"`
	IsActive  bool      `json:"is_active"`
	IsOwner   bool      `gorm:"is_owner" json:"is_owner"`
	CompanyId *int      `gorm:"column:company_id" json:"company_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	Company *Company `gorm:"foreignKey:CompanyId;" json:"company"`
}
