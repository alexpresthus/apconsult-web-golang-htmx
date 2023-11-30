package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// build tailwindcss bundle
//go:generate npm run build

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "public/static")

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	index := func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	}

	e.GET("/", index)

	e.Logger.Fatal(e.Start(":1323"))
}
