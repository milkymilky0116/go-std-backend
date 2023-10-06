package web

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/go-playground/form"
	"github.com/milkymilky0116/go-std-backend/internal/models"
)

type Application struct {
	ServerLogger   *log.Logger
	ErrorLogger    *log.Logger
	Gists          *models.GistModel
	Users          *models.UserModel
	TemplateCache  map[string]*template.Template
	IsTemplateMode bool
	FormDecoder    *form.Decoder
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
	app.ErrorLogger.Println("Not Found")
	app.ClientError(w, http.StatusNotFound)
}

var ErrNoRecords = errors.New("models: no matching record found")
