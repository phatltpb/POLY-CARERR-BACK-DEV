package services

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/addressService"
	authservice "github.com/tuongnguyen1209/poly-career-back/apis/services/authService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/categoryService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/companyService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/jobService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/metaService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/studentJobService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/studentService"
	uploadservice "github.com/tuongnguyen1209/poly-career-back/apis/services/uploadService"
	"github.com/tuongnguyen1209/poly-career-back/apis/services/userService"
)

type Service struct {
	AuthService       authservice.AuthInterface
	AddressService    addressService.AddressServiceInterface
	StudentService    studentService.StudentService
	UserService       userService.UserServiceInterface
	CategoryService   categoryService.CategoryServiceInterface
	JobService        jobService.JobServiceInterface
	CompanyService    companyService.CompanyServiceInterface
	StudentJobService studentJobService.StudentJobServiceInterface
	//
	MetaService   metaService.MetaServiceInterface
	UploadService uploadservice.UploadServiceInterface
}

func Init(repositories *repositories.Repositories) *Service {
	return &Service{
		AuthService:       authservice.Init(repositories),
		AddressService:    addressService.Init(repositories),
		StudentService:    *studentService.Init(repositories),
		UserService:       userService.Init(repositories),
		CategoryService:   categoryService.Init(repositories),
		JobService:        jobService.Init(repositories),
		CompanyService:    companyService.Init(repositories),
		StudentJobService: studentJobService.Init(repositories),

		//
		MetaService:   metaService.Init(repositories),
		UploadService: uploadservice.Init(repositories),
	}
}
