package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.GET("/api/demo/health-check", func(c echo.Context) error {
		fmt.Println("Log health-check in demo")
		log.Info("Log health-check in demo")

		return c.String(http.StatusOK, "it's works")
	})
	e.GET("/api/demo/hello-world", func(c echo.Context) error {
		fmt.Println("Log api-demo-hello-world")
		log.Info("Log api-demo-hello-world")

		return c.String(http.StatusOK, "C4c7us Api Demo")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
