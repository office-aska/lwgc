package app

import (
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
)

type Template struct {
	templates *template.Template
}

var base = "app/views/"

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	ctx := c.Request().Context()
	d, ok := data.(map[string]interface{})
	if ok {
		d["PROJECT_ID"] = environ.GetProjectID()
		user, ok := ctx.Value(model.User{}).(*model.User)
		if ok {
			d["signInUser"] = user
		}
	}
	tmpl := template.Must(t.templates.ParseFiles(base+"layout.html", base+name))
	return tmpl.ExecuteTemplate(w, "layout", data)
}

func NewTemplate() *Template {
	tmpl := template.New("")
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	funcMap := template.FuncMap{
		"ForURI": func(req *http.Request, key string, value string) string {
			params := req.URL.Query()
			params.Set(key, value)
			if key != "page" {
				params.Del("page")
			}
			return "?" + params.Encode()
		},
		"DateTime": func(t time.Time) string {
			if t.IsZero() {
				return ""
			}
			return t.In(jst).Format("2006/01/02 15:04")
		},
		"Format": func(f string, t time.Time) string {
			if t.IsZero() {
				return ""
			}
			return t.In(jst).Format(f)
		},
	}
	tmpl = tmpl.Funcs(funcMap)
	t := &Template{
		templates: template.Must(tmpl.ParseGlob(base + "includes/*.html")),
	}
	return t
}
