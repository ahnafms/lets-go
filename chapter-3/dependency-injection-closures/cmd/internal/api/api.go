package api

import (
	"html/template"
	"net/http"

	"github.com/ahnafms/learn-go/cmd/internal/config"
)

type ApiHandler interface {
	Home(app *config.Application) http.HandlerFunc
}

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go")

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)

		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
