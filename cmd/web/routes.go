package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is main page.")
}

func (app *Application) ViewOneGists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			app.ServerError(w, err)
		}
		gist := findGist(id)
		if gist != (Gist{}) {
			marshaledGist, err := json.Marshal(gist)
			if err != nil {
				app.ServerError(w, err)
			}
			fmt.Fprint(w, string(marshaledGist))
		} else {
			app.NotFound(w)
		}
	}

}

func (app *Application) ViewAllGists(w http.ResponseWriter, r *http.Request) {
	gists := listGist()
	marshaledGistList, err := json.Marshal(gists)
	if err != nil {
		app.ServerError(w, err)
	}
	fmt.Fprint(w, string(marshaledGistList))
}

func (app *Application) CreateGist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Create gist....")
}

func (app *Application) InitRoutes(mux *mux.Router) *mux.Router {

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/gists/view", app.ViewAllGists)
	mux.HandleFunc("/gists/view/{id:[0-9]+}", app.ViewOneGists)
	mux.HandleFunc("/gists/create", app.CreateGist)

	return mux
}
