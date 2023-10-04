package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/milkymilky0116/go-std-backend/cmd/web"
)

func main() {
	mux := mux.NewRouter()
	mux.Use(web.ContentTypeHeader)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &web.Application{
		ServerLogger: infoLog,
		ErrorLogger:  errorLog,
	}

	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  app.InitRoutes(mux),
	}
	infoLog.Println("Server Starting on :4000")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
