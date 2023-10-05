package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/milkymilky0116/go-std-backend/internal/models"
)

func (app *Application) ApiHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/api/gists/view", http.StatusPermanentRedirect)
}

func (app *Application) ApiViewOneGists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			app.NotFound(w)
			return
		}
		gist, err := app.Gists.FindOne(id)
		if err != nil {
			if errors.Is(err, ErrNoRecords) {
				app.NotFound(w)

			} else {
				app.ServerError(w, err)
			}
			return
		}
		marshaledGist, err := json.Marshal(gist)
		if err != nil {
			app.ServerError(w, err)
		}
		fmt.Fprint(w, string(marshaledGist))
	}

}

func (app *Application) ApiViewLatestGists(w http.ResponseWriter, r *http.Request) {
	gists, err := app.Gists.FindMany(10)
	if err != nil {
		app.ServerError(w, err)
	}
	marshaledGistList, err := json.Marshal(gists)
	if err != nil {
		app.ServerError(w, err)
	}
	fmt.Fprint(w, string(marshaledGistList))
}

func (app *Application) ApiCreateGist(w http.ResponseWriter, r *http.Request) {
	var gistParams models.GistParam
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
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

	id, err := app.Gists.Insert(gistParams)
	if err != nil {
		app.ServerError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/gists/view/%d", id), http.StatusSeeOther)
}
