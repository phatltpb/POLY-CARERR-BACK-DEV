package helper

import (
	"github.com/labstack/echo/v4"
	pkg "github.com/tuongnguyen1209/poly-career-back/pkg/jwt"
)

const USER_INFO = "USER_INFO"

func GetJwtFromContext(c echo.Context) *pkg.JwtToken {
	info, ok := c.Get(USER_INFO).(*pkg.JwtToken)
	if !ok {
		return nil
	}
	return info
}
