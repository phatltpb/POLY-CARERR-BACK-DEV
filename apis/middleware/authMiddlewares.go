package middleware

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
	mystatus "github.com/tuongnguyen1209/poly-career-back/pkg/myStatus"
	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
)

func CheckAuth(action string) func(next echo.HandlerFunc) echo.HandlerFunc {
	handleFunc := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			token = strings.TrimSpace(strings.TrimPrefix(token, " "))

			if token == "" {
				token = c.QueryParam("token")
			}

			jwt := pkg.JwtConfig{}

			if token == "" {
				return response.Error(c, &response.CustomError{
					Error: errors.New(validationMessage.TokenMissing),
					Code:  mystatus.TokenMissing,
				})
			}

			claims, err := jwt.Decode(token)
			if err != nil {
				return response.Error(c, &response.CustomError{
					Error: errors.New(validationMessage.TokenInvalid),
					Code:  mystatus.TokenInvalid,
				})
			}
			if claims == nil {
				return response.Error(c, &response.CustomError{
					Error: errors.New(validationMessage.TokenInvalid),
					Code:  mystatus.TokenInvalid,
				})
			}

			if claims.Action != action {
				return response.Error(c, &response.CustomError{
					Error: errors.New(validationMessage.TokenInvalid),
					Code:  mystatus.TokenInvalid,
				})
			}

			c.Set("USER_INFO", claims)
			return next(c)
		}
	}

	return handleFunc
}
