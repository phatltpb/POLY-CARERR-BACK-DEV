package categoryRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type CategoryRepositoryInterface interface {
	GetAll() ([]models.Category, error)
	CreateCate(cate *models.Category) (*models.Category, error)
	GetById(id int) (*models.Category, error)
	UpdateCate(id int, cate *models.Category) error
	DeleteCate(id int) error
}
