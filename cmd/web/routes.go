package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var gists = GistsList()

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is main page.")
}

func ViewOneGists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Fatal(err)
		}
		for _, gist := range gists {
			if gist.Id == id {
				marshaledGist, err := json.Marshal(gist)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprint(w, string(marshaledGist))
				return
			}
		}
		http.Error(w, "Gist Not Found", http.StatusNotFound)
	}

}

func ViewAllGists(w http.ResponseWriter, r *http.Request) {
	gists, err := json.Marshal(gists)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(gists))
}

func CreateGist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Add("Allow", http.MethodPost)
		log.Println("Only POST Method Allowed.")
		return
	}
	fmt.Fprint(w, "Create gist....")
}

func InitRoutes(mux *mux.Router) *mux.Router {

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/gists/view", ViewAllGists)
	mux.HandleFunc("/gists/view/{id:[0-9]+}", ViewOneGists)
	mux.HandleFunc("/gists/create", CreateGist)

	return mux
}
