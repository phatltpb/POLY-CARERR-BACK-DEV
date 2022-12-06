package metaService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/metaDataRepository"
	"github.com/tuongnguyen1209/poly-career-back/pkg/helper"
)

type MetaService struct {
	metaRepository metaDataRepository.MetaHelperInterface
}

func Init(repositories *repositories.Repositories) *MetaService {
	return &MetaService{
		metaRepository: repositories.MetaHelper,
	}
}

func (r *MetaService) GetMeta(modal interface{}, page, limit, total int) *helper.Meta {
	if total < 0 {
		total = r.metaRepository.GetMaxPage(modal)
	}

	return &helper.Meta{
		Page:  page,
		Limit: limit,
		Total: total,
	}
}
