package addressRepository

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (rp *AddressRepository) GetAll() ([]models.City, error) {

	address := []models.City{}

	rs := rp.db.Find(&address)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return address, nil
}

func (rp *AddressRepository) GetById(id int) (*models.City, error) {
	address := &models.City{}

	rs := rp.db.Where("id=?", id).First(address)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return address, nil
}
