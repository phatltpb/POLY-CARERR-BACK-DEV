package helper

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DEFAULT_LIMIT = 10
	DEFAULT_PAGE  = 1
)

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func Paginate(ctx echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.QueryParam("page"))
		if page == 0 {
			page = DEFAULT_PAGE
		}

		limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

		if limit <= 0 {
			limit = DEFAULT_LIMIT
		}

		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}

func Total(model interface{}, db *gorm.DB) int64 {
	var total int64

	if err := db.Model(model).Count(&total).Error; err != nil {
		return 0
	}

	return total
}
