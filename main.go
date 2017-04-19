package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/jinzhu/gorm"
)

func init() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Static("/static", "web/static")
	e.File("/", "web/index.html")

	// Route => handler
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": true})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
