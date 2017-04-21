package db

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
)

func EchoMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("gorm", db)
			next(c)
			return nil
		}
	}
}