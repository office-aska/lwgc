package app

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
)

type Template struct {
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
		csrf := c.Get("csrf")
		if csrf != nil {
			csrfToken := csrf.(string)
			d["csrfField"] = template.HTML(fmt.Sprintf(`<input type="hidden" name="csrf_token" value="%s">`, csrfToken))
			d["csrfToken"] = csrfToken
		}
	}
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
			return t.In(environ.LocationJST()).Format("2006/01/02 15:04")
		},
		"Format": func(f string, t time.Time) string {
			if t.IsZero() {
				return ""
			}
			return t.In(environ.LocationJST()).Format(f)
		},
		"TimeParse": func(s string) time.Time {
			t, err := time.ParseInLocation(time.RFC3339, s, environ.LocationJST())
			if err != nil {
				return time.Time{}
			}
			return t
		},
	}
	tmpl := template.New("")
	tmpl = tmpl.Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseGlob(base + "includes/*.html"))
	tmpl = template.Must(tmpl.ParseFiles(base+"layout.html", base+name))
	return tmpl.ExecuteTemplate(w, "layout", data)
}

func NewTemplate() *Template {
	t := &Template{}
	return t
}
