package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tmp *template.Template

func init() {
	t := template.Must(template.ParseGlob("./ui/html/*.tmpl.html"))
	t = template.Must(t.ParseGlob("./ui/html/**/*.tmpl.html"))
	if t != nil {
		tmp = t
	}
}
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	err := tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) ViewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with id : %d", id)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Creating new Snippet...")
}
