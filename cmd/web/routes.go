package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) InitRoutes() *mux.Router {
	mux := mux.NewRouter()
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	if app.IsTemplateMode {
		mux.HandleFunc("/", app.TemplateHome)
		mux.HandleFunc("/gists/view/{id:[0-9]+}", app.TemplateViewOneGists)
		mux.HandleFunc("/gists/create", app.TemplateCreateGist)
	} else {
		s := mux.PathPrefix("/api").Subrouter()
		s.HandleFunc("/", app.ApiHome)
		s.HandleFunc("/gists/view", app.ApiViewLatestGists)
		s.HandleFunc("/gists/view/{id:[0-9]+}", app.ApiViewOneGists)
		s.HandleFunc("/gists/create", app.ApiCreateGist)
	}
	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NotFound(w)
	})
	mux.Use(app.RecoverPanics)
	mux.Use(app.HTTPLogger)
	mux.Use(app.ContentTypeHeader)
	mux.Use(app.SecureHeaders)
	return mux
}
