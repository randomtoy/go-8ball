package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/randomtoy/go-8ball/pkg/router"
)

func main() {

	router := router.NewRouter()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/", router.MainHandler)
	e.Start(":" + os.Getenv("PORT"))

}
