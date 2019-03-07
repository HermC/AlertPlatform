package router

import (
	"conf"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitContextRouters(e *echo.Echo) error {
	t := &Template{
		templates: template.Must(template.ParseGlob(conf.Conf.Tmpl.Dir + "*" + conf.Conf.Tmpl.Suffix)),
	}
	e.Renderer = t
	
	e.GET("/hello", func(context echo.Context) error {
		return context.Render(http.StatusOK, "hello", "world")
	})
	
	return nil
}
