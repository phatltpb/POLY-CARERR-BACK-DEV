package categoryService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type CategoryServiceInterface interface {
	GetAll() ([]models.Category, *response.CustomError)
	GetById(id int) (*models.Category, *response.CustomError)
	Create(cate *models.Category, user_id int) (*models.Category, *response.CustomError)
	UpdateCate(id int, cate *models.Category, user_id int) *response.CustomError
	Delete(id int, user_id int) *response.CustomError
}
