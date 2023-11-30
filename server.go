package main

import (
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

func goEnvVar(key string) (string, error) {
	godotenv.Load(".env")
	val := os.Getenv(key)
	if val == "" {
		return "", errors.New("Missing env var $" + key)
	}
	return val, nil
}

func main() {
	port, err := goEnvVar("PORT")

	if err != nil {
		log.Fatal(err)
	}
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

	e.Logger.Fatal(e.Start(":" + port))
}
