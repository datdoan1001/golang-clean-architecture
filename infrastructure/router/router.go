package router

import (
	"github.com/garaujo/golang-clean-architecture/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter function stores the required routes for the web application
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })

	return e
}
