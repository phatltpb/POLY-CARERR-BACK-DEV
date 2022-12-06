package addressRepository

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type AddressRepoInterface interface {
	GetAll() ([]models.City, error)
	GetById(id int) (*models.City, error)
}
