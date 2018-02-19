package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", requestHandler)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if isInvalidHost(r) {
		http.Redirect(w, r, "https://check-ip.ml", 301)
		return
	}

	if isInvalidURL(r) {
		http.NotFound(w, r)
		return
	}

	w = addHeaders(w)

	if strings.Contains(r.Host, "raw") {
		serveRawIP(w, r)
	} else {
		serveHTML(w, r)
	}
}

func isInvalidHost(r *http.Request) bool {
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

func isInvalidURL(r *http.Request) bool {
	return r.URL.Path != "/"
}

func addHeaders(w http.ResponseWriter) http.ResponseWriter {
	headers := map[string]string{
		"Strict-Transport-Security": "max-age=31536000; includeSubDomains; preload",
		"X-Frame-Options":           "DENY",
		"X-XSS-Protection":          "1; mode=block",
		"X-Content-Type-Options":    "nosniff",
		"Content-Security-Policy":   "default-src 'none'",
		"Referrer-Policy":           "no-referrer",
		"X-Robots-Tag":              "noarchive",
		"Cache-Control":             "max-age=0",
	}

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	return w
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
