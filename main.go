package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/milkymilky0116/go-std-backend/cmd/cli"
	"github.com/milkymilky0116/go-std-backend/cmd/db"
	"github.com/milkymilky0116/go-std-backend/cmd/web"
)

func main() {

	mux := mux.NewRouter()
	blue := color.New(color.BgBlue).SprintFunc()
	red := color.New(color.BgRed).SprintFunc()

	infoLog := log.New(os.Stdout, blue("INFO\t"), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, red("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
	c, err := cli.ParseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		infoLog.Println(err)
		os.Exit(1)
	}
	db, err := db.DBinit(c)
	if err != nil {
		errorLog.Println(err)
		os.Exit(1)
	}
	infoLog.Println("DB Connected!")
	defer db.Close()
	app := &web.Application{
		ServerLogger: infoLog,
		ErrorLogger:  errorLog,
		DB:           db,
	}
	mux.Use(app.ContentTypeHeader)
	mux.Use(app.HTTPLogger)
	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  app.InitRoutes(mux),
	}
	infoLog.Println("Server Starting on :4000")

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
