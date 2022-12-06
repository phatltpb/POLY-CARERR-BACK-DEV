package categoryService

import (
	"errors"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/categoryRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/userRepository"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

type CategoryService struct {
	categoryRepository categoryRepository.CategoryRepositoryInterface
	userRepository     userRepository.UserRepositoryInterface
}

func Init(repositories *repositories.Repositories) *CategoryService {
	return &CategoryService{
		categoryRepository: repositories.CategoryRepository,
		userRepository:     repositories.UserRepository,
	}
}

func (s *CategoryService) GetAll() ([]models.Category, *response.CustomError) {

	categories, err := s.categoryRepository.GetAll()
	if err != nil {
		return nil, &response.CustomError{
			Error: err,
		}
	}

	return categories, nil
}

func (s *CategoryService) Create(cate *models.Category, user_id int) (*models.Category, *response.CustomError) {
	if cate == nil {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	if !s.userRepository.CheckingUserAdmin(user_id) {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.Unauthorized),
		}
	}

	create, err := s.categoryRepository.CreateCate(cate)
	if err != nil {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return create, nil
}

func (s *CategoryService) GetById(id int) (*models.Category, *response.CustomError) {
	if id <= 0 {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	cate, err := s.categoryRepository.GetById(id)
	if err != nil {
		return nil, &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return cate, nil
}

func (s *CategoryService) UpdateCate(id int, cate *models.Category, user_id int) *response.CustomError {
	if id <= 0 {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	if !s.userRepository.CheckingUserAdmin(user_id) {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.Unauthorized),
		}
	}
	if err := s.categoryRepository.UpdateCate(id, cate); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return nil
}

func (s *CategoryService) Delete(id int, user_id int) *response.CustomError {
	if id <= 0 {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.IsValid),
		}
	}
	if !s.userRepository.CheckingUserAdmin(user_id) {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.Unauthorized),
		}
	}
	if err := s.categoryRepository.DeleteCate(id); err != nil {
		return &response.CustomError{
			Code:  mystatus.BadRequest,
			Error: errors.New(validationMessage.HaveSomeErr),
		}
	}
	return nil
}
