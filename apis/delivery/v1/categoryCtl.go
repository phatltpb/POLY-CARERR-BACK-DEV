package v1

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"github.com/tuongnguyen1209/poly-career-back/apis/services"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
)

type CategoryCtl struct {
	services *services.Service
}

func InitCategoryCtl(services *services.Service) *CategoryCtl {
	return &CategoryCtl{
		services: services,
	}
}

func (ctl *CategoryCtl) GetAll(c echo.Context) error {
	categories, err := ctl.services.CategoryService.GetAll()

	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, categories)
}

func (ctl *CategoryCtl) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	cate, err1 := ctl.services.CategoryService.GetById(id)
	if err1 != nil {
		return response.Error(c, err1)
	}
	return response.Success(c, cate)
}

func (ctl *CategoryCtl) CreateCate(c echo.Context) error {
	var cate models.Category
	if err := c.Bind(&cate); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	createCate, err1 := ctl.services.CategoryService.Create(&cate, jwtToken.Id)
	if err1 != nil {
		return response.Error(c, err1)
	}

	return response.Success(c, createCate)
}

func (ctl *CategoryCtl) UpdateCate(c echo.Context) error {
	var cate models.Category
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	if err := c.Bind(&cate); err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	if err := ctl.services.CategoryService.UpdateCate(id, &cate, jwtToken.Id); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}

func (ctl *CategoryCtl) DeleteCate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.Error(c, &response.CustomError{
			Error: err,
		})
	}
	jwtToken := c.Get(USER_INFO).(*pkg.JwtToken)
	if err := ctl.services.CategoryService.Delete(id, jwtToken.Id); err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, "ok")
}
