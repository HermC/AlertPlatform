package router

import (
	. "../conf"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	hosts := make(map[string]*Host)

	hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}
	hosts[Conf.Server.DomainApi] = &Host{api.Routers()}
	hosts[Conf.Server.DomainSocket] = &Host{socket.Routers()}

	return hosts
}

func RunSubdomains(configFilePath string) {
	if err := InitConfig(configFilePath); err != nil {
		log.Panic(err)
	}
	log.SetLevel(GetLogLvl())

	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())
	e.Logger.SetLevel(GetLogLvl())

	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	e.Use(mw.MethodOverride())

	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string {"*"},
		AllowHeaders: []string {echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))
	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error: %v", _err)
		}

		host := hosts[u.Hostname()]
		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})
}