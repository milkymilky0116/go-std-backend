package web

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/form"
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

func (app *Application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	err = app.FormDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalideDecodeError *form.InvalidDecoderError
		if errors.As(err, &invalideDecodeError) {
			panic(err)
		}
		return err
	}
	return nil
}
