package addressService

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/addressRepository"
)

type AddressService struct {
	addressRepository addressRepository.AddressRepoInterface
}

func Init(repository *repositories.Repositories) *AddressService {
	return &AddressService{
		addressRepository: repository.AddressRepository,
	}
}

func (s *AddressService) GetAddress() ([]models.City, error) {
	rs, err := s.addressRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return rs, nil
}
