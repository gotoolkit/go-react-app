package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gotoolkit/db/config"
	"github.com/gotoolkit/db/orm"
	"github.com/gotoolkit/db/handler"
	"github.com/gotoolkit/db/model"
	"github.com/gotoolkit/db"
)

// User simple model
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}


func main() {

	dbConfig := &config.DBConfig{
		Dialect:  "mysql",
		Username: "root",
		Password: "root",
		Host:     "dockerhost",
		Port:     "3306",
		Name:     "go-react-db",
		Charset:  "utf8",
	}

	orm.InitialDB(dbConfig)

	orm.GetDB().AutoMigrate(&model.User{}, &model.Task{})

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(db.EchoMiddleware(db))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Static("/static", "web/static")
	e.File("/", "web/index.html")

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": true})
	})

	handler.SetRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}


