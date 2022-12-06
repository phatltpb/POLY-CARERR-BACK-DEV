package addressService

import "github.com/tuongnguyen1209/poly-career-back/apis/models"

type AddressServiceInterface interface {
	GetAddress() ([]models.City, error)
}
