package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return template.Must(t.templates[name], nil).Execute(w, data)
}

func LoadTemplates(e *echo.Echo) {
	m := make(map[string]*template.Template)

	var err error

	// Add templates here:
	m["main"], err = template.ParseFiles("templates/base.html", "templates/main.html")
	m["search"], err = template.ParseFiles("templates/base.html", "templates/search.html")

	if err != nil {
		fmt.Println("[ERR] Failed to load templates:\n", err)
	}

	t := &Template{m}
	e.Renderer = t
}
