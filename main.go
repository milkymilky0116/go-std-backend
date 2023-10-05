package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/milkymilky0116/go-std-backend/cmd/cli"
	"github.com/milkymilky0116/go-std-backend/cmd/db"
	"github.com/milkymilky0116/go-std-backend/cmd/web"
	"github.com/milkymilky0116/go-std-backend/internal/models"
)

func main() {

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
		ServerLogger:   infoLog,
		ErrorLogger:    errorLog,
		Gists:          &models.GistModel{DB: db},
		Users:          &models.UserModel{DB: db},
		IsTemplateMode: c.IsTemplateMode,
	}

	mux := app.InitRoutes()
	defaultMode := "API Mode"
	if app.IsTemplateMode {
		defaultMode = "Template Mode"
		templateCache, err := web.NewTemplateCache()
		if err != nil {
			errorLog.Fatal(err)
		}
		app.TemplateCache = templateCache
	}

	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  mux,
	}

	startLog := fmt.Sprintf("Server Starting on :4000 / %s", defaultMode)
	infoLog.Println(startLog)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
