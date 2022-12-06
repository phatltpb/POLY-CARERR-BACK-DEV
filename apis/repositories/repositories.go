package repositories

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/addressRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/applyJobRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/categoryRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/companyRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/jobRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/metaDataRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentCvRepository"
	studentRepository "github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/studentWatchRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/uploadRepository"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories/userRepository"
	"gorm.io/gorm"
)

type Repositories struct {
	StudentRepository  studentRepository.StudentRepositoryInterface
	AddressRepository  addressRepository.AddressRepoInterface
	UserRepository     userRepository.UserRepositoryInterface
	CategoryRepository categoryRepository.CategoryRepositoryInterface
	JobRepository      jobRepository.JobRepositoryInterface
	CompanyRepository  companyRepository.CompanyRepositoryInterface
	UploadRepository   uploadRepository.UpLoadRepositoryInterface
	ApplyJobRepository applyJobRepository.ApplyJobRepositoryInterface
	StudentCv          studentCvRepository.StudentCvRepositoryInterface
	StudentWatch       studentWatchRepository.StudentWatchRepoInterface
	// meta
	MetaHelper metaDataRepository.MetaHelperInterface
}

func Init(db *gorm.DB) *Repositories {
	return &Repositories{
		StudentRepository:  studentRepository.Init(db),
		AddressRepository:  addressRepository.Init(db),
		UserRepository:     userRepository.Init(db),
		CategoryRepository: categoryRepository.Init(db),
		JobRepository:      jobRepository.Init(db),
		CompanyRepository:  companyRepository.Init(db),
		UploadRepository:   uploadRepository.Init(),
		ApplyJobRepository: applyJobRepository.Init(db),
		StudentCv:          studentCvRepository.Init(db),
		StudentWatch:       studentWatchRepository.Init(db),
		// meta
		MetaHelper: metaDataRepository.Init(db),
	}
}
