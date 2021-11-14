package inf

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
	templates := &Template{make(map[string]*template.Template)}

	templates.load("main")
	templates.load("error")
	templates.load("search")
	templates.load("register")
	templates.load("login")
	templates.load("userlist")
	templates.load("register_succ")
	templates.load("item")
	templates.load("add_item")

	e.Renderer = templates
}

func (t *Template) load(name string) {
	var err error
	t.templates[name], err = template.ParseFiles("data/templates/base.html", "data/templates/"+name+".html")
	if err != nil {
		fmt.Printf("[ERROR] Failed to load template '%v'.\n", name)
	}
}
