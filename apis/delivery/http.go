package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tuongnguyen1209/poly-career-back/apis/controller"
	v1 "github.com/tuongnguyen1209/poly-career-back/apis/delivery/v1"
	"github.com/tuongnguyen1209/poly-career-back/apis/repositories"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	group := e.Group("/api")

	repositories := repositories.Init(db)

	services := services.Init(repositories)

	v1.Init(*group, services)

	e.RouteNotFound("/*", controller.HandleNotFound)

	e.HTTPErrorHandler = e.DefaultHTTPErrorHandler

	return e
}
