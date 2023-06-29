package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CheckToken(ctx echo.Context) (*jwt.Token, error) {
	bearer_token := ctx.Request().Header.Get("Authorization")

	if len(bearer_token) <= 0 {
		return nil, ctx.JSON(http.StatusUnauthorized, "Authorization Header Not Found")
	}

	token := strings.Split(bearer_token, " ")
	if token[1] == "" {
		ctx.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	// Check - Parse token
	parsedToken, err := ValidateToken(token[1])

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Invalid Token",
		})
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
	}

	return parsedToken, nil
}

func GetPayload(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token Claims")
	}
	return claims, nil
}

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token, _ := CheckToken(ctx)

			claims, _ := GetPayload(token)

			// Check Role admin
			role, ok := claims["Role"].(string)
			if !ok || role != "R1" {
				ctx.JSON(http.StatusUnauthorized, echo.Map{
					"Message": "Unauthorized",
				})
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			return next(ctx)
		}
	}
}

func IsManagement() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token, _ := CheckToken(ctx)

			claims, _ := GetPayload(token)

			// Check Role manager
			role, ok := claims["Role"].(string)
			if !ok || role != "R2" {
				ctx.JSON(http.StatusUnauthorized, echo.Map{
					"Message": "Unauthorized",
				})
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			return next(ctx)
		}
	}
}

func IsUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token, _ := CheckToken(ctx)

			claims, _ := GetPayload(token)

			// Check Role user
			role, ok := claims["Role"].(string)
			if !ok || role != "R3" {
				ctx.JSON(http.StatusUnauthorized, echo.Map{
					"Message": "Unauthorized",
				})
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			return next(ctx)
		}
	}
}
