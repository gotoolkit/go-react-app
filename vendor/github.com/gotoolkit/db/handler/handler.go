package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func SetModelRoute(e *echo.Echo, path string, in interface{}) {

	e.GET(path, GetAll(in))
	e.GET(fmt.Sprint(path, "/:id"), GetByID(in))
	e.POST(path, Create(in))
}

func Create(writer interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		orm := c.Get("orm").(*gorm.DB)
		v := reflect.TypeOf(writer).Elem()
		m := reflect.New(v).Interface()
		err := c.Bind(m)
		if err != nil {
			return err
		}
		return process(c, orm.Save(m).Error, m)
	}
}

func GetByID(reader interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		orm := c.Get("orm").(*gorm.DB)

		v := reflect.TypeOf(reader).Elem()
		m := reflect.New(v).Interface()

		err := orm.Find(m, id).Error
		return process(c, err, m)
	}
}
func GetAll(reader interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		orm := c.Get("orm").(*gorm.DB)
		v := reflect.TypeOf(reader).Elem()
		m := reflect.New(reflect.SliceOf(v)).Interface()
		err := orm.Find(m).Error
		return process(c, err, m)
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
