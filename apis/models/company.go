package models

type Company struct {
	Id                int              `gorm:"primary_key;column:id" json:"id"`
	Name              string           `gorm:"column:name" json:"name"`
	LocationId        int              `gorm:"column:location" json:"province_id"`
	Location          *City            `gorm:"foreignKey:LocationId" json:"province"`
	Address           string           `gorm:"column:address" json:"address"`
	Avatar            string           `gorm:"column:avatar" json:"avatar"`
	Banner            string           `gorm:"column:banner" json:"banner"`
	TaxCode           string           `gorm:"column:tax_code" json:"tax_code"`
	Website           string           `gorm:"column:website" json:"website"`
	Phone             string           `gorm:"column:phone" json:"phone"`
	Information       string           `gorm:"column:information" json:"information"`
	CompanyActivityId int              `gorm:"column:company_activity" json:"company_activity_id"`
	CompanyActivity   *CompanyActivity `gorm:"foreignKey:CompanyActivityId" json:"company_activity"`
	Size              int              `gorm:"column:size" json:"size"`
	Status            int              `gorm:"column:status" json:"status"`
	Hidden            bool             `gorm:"column:is_hidden" json:"is_hidden"`
	Active            bool             `gorm:"column:is_active" json:"is_active"`
	Noti              bool             `gorm:"column:is_noti" json:"is_noti"`
}

type CompanyActivity struct {
	Id   int    `gorm:"primary_key;column:id" json:"id"`
	Name string `json:"name"`
}
