package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func AuthorizeJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			const BearerSchema string = "Bearer "
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "No Authorization header found",
				})
			}

			tokenString := authHeader[len(BearerSchema):]

			if token, err := ValidateToken(tokenString); err != nil {
				fmt.Println("token", tokenString, err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Not Valid Token",
				})
			} else {
				if claims, ok := token.Claims.(jwt.MapClaims); !ok {
					return c.NoContent(http.StatusUnauthorized)
				} else {
					if token.Valid {
						c.Set("userID", claims["userID"])
						fmt.Println("during authorization", claims["userID"])
					} else {
						return c.NoContent(http.StatusUnauthorized)
					}
				}
			}

			return next(c)
		}
	}
}
