package server

import (
	"echo-get-started/di"
	"echo-get-started/rest/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	systemsRouter(e)
	nounRouter(e)
	e.Logger.Fatal(e.Start(":8088"))
	return e
}

func systemsRouter(e *echo.Echo) {
	e.GET("/v1/systems/ping", handler.Ping())
}

func nounRouter(e *echo.Echo) {
	mainGroupHandler := di.InitNounHandler()
	v1 := e.Group("/nouns")
	v1.GET("", mainGroupHandler.GetCountByNoun())
}
