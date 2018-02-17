package main

import (
    "net/http"

    "google.golang.org/appengine"
    "html/template"
)

var (
    indexTemplate = template.Must(template.ParseFiles("index.html"))
)

type templateParams struct {
    IP string
}

func main() {
    http.HandleFunc("/", indexHandler)
    appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    params := templateParams{IP: r.RemoteAddr}

    indexTemplate.Execute(w, params)
}
