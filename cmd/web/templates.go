package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/milkymilky0116/go-std-backend/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("2006 Jan 02 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.gohtml")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.gohtml")

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
