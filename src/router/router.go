package router

import (
	"conf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouters() error {
	e := echo.New()

	e.Use(middleware.CORS())
	//e.GET("/", func(context echo.Context) error {
	//	return context.String(http.StatusOK, "hello")
	//})

	_ = InitContextRouters(e)

	e.Logger.SetLevel(conf.GetLogLvl())
	e.Logger.Info(e.Start(conf.Conf.Server.Addr))



	return nil
}