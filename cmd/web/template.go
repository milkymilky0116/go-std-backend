package web

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func humanDate(t time.Time) string {
	return t.Format("2006 Jan 02 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
func (app *Application) TemplateHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}
	data := app.newTemplateData(r)
	gists, err := app.Gists.FindMany(10)
	if err != nil {
		app.ServerError(w, err)
	}
	data.Gists = gists
	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *Application) TemplateViewOneGists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			app.ServerError(w, err)
		}
		data := app.newTemplateData(r)
		gist, err := app.Gists.FindOne(id)
		if err != nil {
			if errors.Is(err, ErrNoRecords) {
				app.NotFound(w)
			} else {
				app.ServerError(w, err)
			}
			return
		}
		data.Gist = gist
		app.render(w, http.StatusOK, "view.tmpl.html", data)
	}
}

func (app *Application) TemplateViewAllGists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "View all Gist")
}

func (app *Application) TemplateCreateGist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create gist...")
}
