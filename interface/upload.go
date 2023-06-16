package _interface

import (
	"github.com/labstack/echo/v4"
)

type UploadFile interface {
	Upload(ctx echo.Context) string
}

type Upload struct {
	it UploadFile
}
