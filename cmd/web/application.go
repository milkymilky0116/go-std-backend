package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type Application struct {
	ServerLogger *log.Logger
	ErrorLogger  *log.Logger
	DB           *sql.DB
}

func (app *Application) Logger(w http.ResponseWriter, context string) {
	app.ServerLogger.Output(1, context)
}
func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
