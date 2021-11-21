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
	if _, ok := t.templates[name]; !ok {
		return NoTemplateError(name)
	}

	return template.Must(t.templates[name], nil).Execute(w, data)
}

func LoadTemplates(e *echo.Echo) {
	templates := &Template{make(map[string]*template.Template)}

	templates.load("main")
	templates.load("error")
	templates.load("search")
	templates.load("register")
	templates.load("login")
	templates.load("user_list")
	templates.load("status")
	templates.load("item")
	templates.load("add_item")
	templates.load("confirm")

	e.Renderer = templates
}

func (t *Template) load(name string) {
	var err error
	t.templates[name], err = template.ParseFiles("data/templates/base.html", "data/templates/"+name+".html")

	if err != nil {
		fmt.Printf("[ERROR] Failed to load template '%v':\n%v", name, err)
	}
}

type NoTemplateError string

func (e NoTemplateError) Error() string {
	return fmt.Sprintf("The template '%v' has not been registered.", string(e))
}
