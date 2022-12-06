package v1

import (
	"github.com/tuongnguyen1209/poly-career-back/apis/controller"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
)

const USER_INFO = "USER_INFO"

type Controller struct {
	AuthCtl     *controller.AuthController
	CategoryCtl *CategoryCtl
	JobCtl      *JobCtl
	ProfileCtl  *ProfileController
	AddressCtl  *AddressCtl
	CompanyCtl  *CompanyCtl
	UploadCtl   *UploadCtl
	ApplyJobCtl *ApplyJobCtl
}

func InitController(services *services.Service) *Controller {
	return &Controller{
		AuthCtl:     controller.Init(services),
		CategoryCtl: InitCategoryCtl(services),
		JobCtl:      InitJobCtl(services),
		ProfileCtl:  InitProfileCtl(services),
		AddressCtl:  InitAddressCtl(services),
		CompanyCtl:  InitCompanyCtl(services),
		UploadCtl:   InitUpload(services),
		ApplyJobCtl: InitApplyJobCtl(services),
	}
}
