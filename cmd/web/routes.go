package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/gists/view", http.StatusPermanentRedirect)
}

func (app *Application) ViewOneGists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			app.ServerError(w, err)
		}
		gist, err := app.findGist(id)
		if err != nil {
			app.ServerError(w, err)
		}
		marshaledGist, err := json.Marshal(gist)
		if err != nil {
			app.ServerError(w, err)
		}
		fmt.Fprint(w, string(marshaledGist))
	}

}

func (app *Application) ViewAllGists(w http.ResponseWriter, r *http.Request) {
	gists, err := app.listGist()
	if err != nil {
		app.ServerError(w, err)
	}
	marshaledGistList, err := json.Marshal(gists)
	if err != nil {
		app.ServerError(w, err)
	}
	fmt.Fprint(w, string(marshaledGistList))
}

func (app *Application) CreateGist(w http.ResponseWriter, r *http.Request) {
	var gistParams GistParam
	if r.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		app.ServerError(w, err)
	}

	err = json.Unmarshal(b, &gistParams)
	if err != nil {
		app.ServerError(w, err)
	}

	newGist, err := app.appendGist(gistParams)
	if err != nil {
		app.ClientError(w, http.StatusMisdirectedRequest)
	}

	marshaledGist, err := json.Marshal(newGist)
	if err != nil {
		app.ServerError(w, err)
	}

	fmt.Fprint(w, string(marshaledGist))
}

func (app *Application) InitRoutes(mux *mux.Router) *mux.Router {

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/gists/view", app.ViewAllGists)
	mux.HandleFunc("/gists/view/{id:[0-9]+}", app.ViewOneGists)
	mux.HandleFunc("/gists/create", app.CreateGist)

	return mux
}
