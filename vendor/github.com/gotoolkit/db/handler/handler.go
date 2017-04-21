package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/gotoolkit/db/model"
	"log"
)

func SetRoute(e *echo.Echo) {

	setModelRoute(e, "/tasks", &model.Task{})
	setModelRoute(e, "/users", &model.User{})
}

func setModelRoute(e *echo.Echo, path string, in model.ReadWriter) {

	e.GET(path, GetAll(in))
	e.GET(fmt.Sprint(path,"/:id"), GetByID(in))
	e.POST(path, Create(in))
}

func Create(writer model.ReadWriter) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := writer.BindModel()
		err := c.Bind(user)
		if err != nil {
			return err
		}
		err = writer.Create(user)
		return process(c, err, user)
	}
}

func GetByID(reader model.Reader) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		log.Println(id)
		result, err := reader.GetByID(id)
		return process(c, err, result)
	}
}
func GetAll(reader model.Reader) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := reader.GetAll()
		return process(c, err, result)
	}
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
