package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"google.golang.org/appengine"
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
	if i := strings.Index(r.Host, ":"); i != -1 {
		return r.Host[:i]
	}

	return r.Host
}

func serveRawIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.RemoteAddr)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("index.html"))

	param := map[string]string{
		"IP": r.RemoteAddr,
	}

	indexTemplate.Execute(w, param)
}
