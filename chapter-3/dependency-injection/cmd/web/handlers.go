package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		app.log.Error(err.Error(), "method", r.Method, r.URL.RequestURI())
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}

func (app *application) downloadFile(w http.ResponseWriter, r *http.Request) {
	saveUrl := filepath.Clean(r.URL.Path)

	log.Print(saveUrl)
	r.URL.Path = saveUrl

	http.ServeFile(w, r, "./ui/static/img/logo.png")
}

func (app *application) downloadFileFromUserInput(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	sanitizedPath := filepath.Clean(path)

	basePath := "./ui/static/img/"

	fullPath := basePath + sanitizedPath

	log.Print(filepath.Clean(fullPath))

	log.Print(fullPath)
	http.ServeFile(w, r, fullPath)
}
