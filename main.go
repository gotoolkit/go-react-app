package main

import (
	"net/http"

	"fmt"
	"github.com/gotoolkit/db/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

var db *gorm.DB
var err error
// User simple model
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func init() {
	dbConfig := &config.DBConfig{
		Dialect:  "mysql",
		Username: "root",
		Password: "root",
		Host:     "dockerhost",
		Port:     "3306",
		Name:     "go-react-db",
		Charset:  "utf8",
	}
	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset,
	)
	db, err = gorm.Open(dbConfig.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
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

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": true})
	})

	e.GET("/users", GetList)

	e.Logger.Fatal(e.Start(":8080"))
}

func GetDB() *gorm.DB {
	return db
}

func GetList(c echo.Context) error {
	users := &[]User{}
	return process(c, GetDB().Find(&users).Error, &users)
}
func process(c echo.Context, err error, result interface{}) error {
	var msg string
	if err != nil {
		msg = fmt.Sprint(err)
	}
	statusCode := http.StatusOK
	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = http.StatusNotFound
	case gorm.ErrInvalidSQL, gorm.ErrInvalidTransaction, gorm.ErrCantStartTransaction:
		statusCode = http.StatusInternalServerError
	}
	return c.JSON(statusCode, echo.Map{
		"resule":  result,
		"error":   err != nil,
		"message": msg,
	})
}
