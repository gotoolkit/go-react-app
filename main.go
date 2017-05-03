package main

import (
	"net/http"
	"strings"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gotoolkit/db/handler"
	//"github.com/gotoolkit/db"
	"github.com/gotoolkit/db/model"
	//"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// User simple model
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	//orm := db.InitDBWithConfig(db.DBConfig{
	//	Dialect:  "postgres",
	//	Username: "postgres",
	//	Password: "D6EaUhzvWKnaMqq",
	//	Host:     "db",
	//	Port:     "3306",
	//	Name:     "go-react-db",
	//	Charset:  "utf8",
	//})
	//
	//orm.AutoMigrate(&model.User{})

	e := echo.New()

	//e.Use(db.EchoMiddleware(orm))

	//ordered
	setupMiddleWare(e)
	setupUI(e)
	setupRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func setupAuth(e *echo.Echo) {
	e.GET("/token", func(c echo.Context) error {

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()


		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	g := e.Group("/admin")
	g.Use(middleware.JWT([]byte("secret")))
	g.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome "+name+"!")
	})
	//g.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
	//	Validator: func(username, password string, c echo.Context) bool {
	//		if username == "paul" && password == "tian" {
	//			return true
	//		}
	//		return false
	//	},
	//}))


}

func setupMiddleWare(e *echo.Echo) {

	// log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasSuffix(c.Request().Host, "l.wizmacau.com") {
				return true
			}
			return false
		},
	}))

	// recover
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	setupAuth(e)
}

func setupUI(e *echo.Echo) {

	e.Static("/static", "web/static")
	e.File("/", "web/index.html")

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": true})
	})
}

func setupRoute(e *echo.Echo) {
	handler.SetModelRoute(e, "/users", &model.User{})
}
