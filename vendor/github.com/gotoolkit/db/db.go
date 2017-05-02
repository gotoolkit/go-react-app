package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"log"
)

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Charset  string
}

var DefaultDBConfig = DBConfig{
	Dialect:  "mysql",
	Username: "guest",
	Password: "password",
	Host:     "localhost",
	Port:     "3306",
	Name:     "todo",
	Charset:  "utf8",
}

func InitDB() *gorm.DB {
	return InitDBWithConfig(DefaultDBConfig)

}

func InitDBWithConfig(dbConfig DBConfig) *gorm.DB {
	configUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset,
	)
	db, err := gorm.Open(dbConfig.Dialect, configUri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func EchoMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("orm", db)
			next(c)
			return nil
		}
	}
}
