package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/middleware"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
)

func Init(group echo.Group, services *services.Service) {
	// router v1
	v1Router := group.Group("/v1")
	controller := InitController(services)
	// auth
	authRouter := v1Router.Group("/auth")
	authRouter.POST("/login", controller.AuthCtl.Login)
	authRouter.POST("/register", controller.AuthCtl.Register)
	authRouter.POST("/admin/login", controller.AuthCtl.LoginAdmin)
	authRouter.GET("/verify-account/:token", controller.AuthCtl.VerifyRegister)
	authRouter.POST("/admin/register", controller.AuthCtl.RegisterAdmin)
	authRouter.GET("/admin/verify-account/:token", controller.AuthCtl.VerifyAdminRegister)

	// reset password
	authRouter.POST("/forgot-password", controller.AuthCtl.ForgotStudentPassword)
	authRouter.POST("/reset-password/:token", controller.AuthCtl.ResetPassword)
	// reset password
	authRouter.POST("/admin/forgot-password", controller.AuthCtl.ForgotAdminPassword)
	authRouter.POST("/admin/reset-password/:token", controller.AuthCtl.ResetAdminPassword)

	// province
	v1Router.GET("/province", controller.AddressCtl.GetAddress)
	// upload file
	v1Router.POST("/upload", controller.UploadCtl.UploadImage)
	v1Router.POST("/upload/file", controller.UploadCtl.UploadFile)
	// companies
	v1Router.GET("/companies", controller.CompanyCtl.GetCompany)
	v1Router.GET("/company/:id", controller.CompanyCtl.GetCompanyById)
	v1Router.GET("/company_activities", controller.CompanyCtl.GetCompanyActivity)
	// jobs
	v1Router.GET("/jobs", controller.JobCtl.GetAllJobs)
	v1Router.GET("/job/fit", controller.JobCtl.GetJobFit)
	v1Router.GET("/job/:id", controller.JobCtl.GetJobById)
	/// categories
	v1Router.GET("/categories", controller.CategoryCtl.GetAll)
	v1Router.GET("/categories/:id", controller.CategoryCtl.GetById)

	v1Router.GET("/verify/change_email/:token", controller.ProfileCtl.VerifyChangeStudentEmail)
	// student router
	studentGroup := v1Router.Group("/student")
	StudentRouter(studentGroup, controller)

	// employer router
	EmployerRouter(v1Router.Group("/"), controller)
	// company
	// EmployerRouter(v1Router.Group("/company"), controller)
}

func StudentRouter(group *echo.Group, controller *Controller) {
	group.Use(middleware.CheckAuth(pkg.JwtAction.Login))
	group.GET("/profile", controller.ProfileCtl.GetMyProfileStudent)
	group.PUT("/profile", controller.ProfileCtl.UpdateProfileStudent)
	group.PUT("change-password", controller.ProfileCtl.ChangePassword)
	// cv
	group.POST("/apply", controller.ApplyJobCtl.CreateApplyJobWithOldCv)
	group.POST("/apply/new", controller.ApplyJobCtl.CreateApplyJobWithNewCv)
	group.GET("/mycv", controller.ApplyJobCtl.GetStudentCv)

	// profile
	group.PUT("/profile/detail", controller.ProfileCtl.UpdateStudentDetail)

	// education
	group.POST("/profile/education", controller.ProfileCtl.CreateStudentEducation)
	group.PUT("/profile/education/:id", controller.ProfileCtl.UpdateStudentEducation)
	group.DELETE("/profile/education/:id", controller.ProfileCtl.DeleteStudentEducation)

	//change email
	group.PUT("/profile/change_email", controller.ProfileCtl.UpdateStudentEmail)
}

func EmployerRouter(group *echo.Group, controller *Controller) {
	group.Use(middleware.CheckAuth(pkg.JwtAction.LoginEmployer))
	group.GET("admin/profile", controller.ProfileCtl.GetMyProfileEmployer)
	group.PUT("admin/profile", controller.ProfileCtl.UpdateMyProfileEmployer)
	// profile
	group.PUT("admin/change-password", controller.ProfileCtl.ChangeAdminPassword)
	group.POST("job", controller.JobCtl.CreateJob)
	group.PUT("job/:id", controller.JobCtl.UpdateJob)
	// company
	group.POST("company", controller.CompanyCtl.CreateCompany)
	group.PUT("company", controller.CompanyCtl.UpdateCompany)
	//apply_job
	group.GET("employer/apply_job", controller.ApplyJobCtl.GetCvByEmployerId)
	group.GET("employer/apply_job/:id", controller.ApplyJobCtl.GetCvById)

	group.PUT("employer/apply_job/:id", controller.ApplyJobCtl.UpdateStatusApplyJob)
	// categories
	group.POST("admin/categories", controller.CategoryCtl.CreateCate)
	group.PUT("admin/categories/:id", controller.CategoryCtl.UpdateCate)
	group.DELETE("admin/categories/:id", controller.CategoryCtl.DeleteCate)
}
