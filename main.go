package main

import (
    "fmt"
    "net/http"
    "google.golang.org/appengine"
    "strings"
    "html/template"
)

func main() {
    http.HandleFunc("/", requestHandler)
    appengine.Main()
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
    if isInvalidURL(r) {
        http.Redirect(w, r, "https://check-ip.ml", 301)
    }

    if strings.Contains(r.Host, "raw") {
        serveRawIP(w, r)
    } else {
        serveHTML(w, r)
    }
}

func isInvalidURL(r *http.Request) bool {
    host := getHost(r)

    return !strings.HasSuffix(host, "check-ip.ml") &&
        !strings.HasSuffix(host, "localhost")
}

func getHost(r *http.Request) string {
    host := r.Host

    if i := strings.Index(r.Host, ":"); i != -1 {
        host = r.Host[:i]
    }

    return host
}

func serveRawIP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, r.RemoteAddr)
}

var (
    indexTemplate = template.Must(template.ParseFiles("index.html"))
)

type templateParams struct {
    IP string
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
    params := templateParams{IP: r.RemoteAddr}

    indexTemplate.Execute(w, params)
}
