package web

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *Application) render(w http.ResponseWriter, status int, page string, data *TemplateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("template %s does not exists", page)
		app.ServerError(w, err)
		return
	}
	buffer := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		app.ServerError(w, err)
	}

	w.WriteHeader(status)
	buffer.WriteTo(w)
}

func (app *Application) newTemplateData(r *http.Request) *TemplateData {
	return &TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
