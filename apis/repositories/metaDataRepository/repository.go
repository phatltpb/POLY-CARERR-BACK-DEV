package metaDataRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
	"gorm.io/gorm"
)

type MetaHelper struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *MetaHelper {
	return &MetaHelper{
		db: db,
	}
}

func (m *MetaHelper) GetMaxPage(model interface{}) int {
	return int(helper.Total(model, m.db))
}
