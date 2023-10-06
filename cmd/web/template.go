package web

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/milkymilky0116/go-std-backend/internal/models"
	"github.com/milkymilky0116/go-std-backend/internal/validator"
)

type gistCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	validator.Validator `form:"-"`
}

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
	data := app.newTemplateData(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			app.ServerError(w, err)
		}

		gist, err := app.Gists.FindOne(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				app.NotFound(w)
			} else {
				app.ServerError(w, err)
			}
			return
		}
		data.Gist = gist
	}
	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

func (app *Application) TemplateViewAllGists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "View all Gist")
}

func (app *Application) TemplateCreateGistGet(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = gistCreateForm{
		Title:   "",
		Content: "",
	}
	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

func (app *Application) TemplateCreateGistPost(w http.ResponseWriter, r *http.Request) {
	current_id, err := strconv.Atoi(r.Header.Get("current_user"))
	if err != nil {
		app.ServerError(w, err)
		return
	}
	var form gistCreateForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank.")
	form.CheckField(validator.MaxChars(form.Title, 128), "title", "This field cannot be more than 100 characters long.")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank.")
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}
	params := &models.GistParam{
		Title:   form.Title,
		Content: form.Content,
		Writer:  current_id,
	}
	id, err := app.Gists.Insert(*params)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/gists/view/%d", id), http.StatusSeeOther)
}
