package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milkymilky0116/go-std-backend/cmd/web"
)

func main() {
	mux := mux.NewRouter()
	mux.Use(web.ContentTypeHeader)
	web.InitRoutes(mux)
	log.Println("Server starting on :4000")
	http.ListenAndServe(":4000", mux)
}
