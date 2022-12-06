package metaService

import "github.com/tuongnguyen1209/poly-career-back/pkg/helper"

type MetaServiceInterface interface {
	GetMeta(modal interface{}, page, limit, total int) *helper.Meta
}
