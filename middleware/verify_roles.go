package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckValidateToken(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
	}

	return nil
}

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			// check role_code
			CheckValidateToken(ctx)

			//payload := jwt.C

			return next(ctx)
		}
	}
}

func IsManagement() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// check role_code

			return next(c)
		}
	}
}

func IsUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// check role_code

			return next(c)
		}
	}
}
