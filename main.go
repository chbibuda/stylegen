package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var styles []string = []string{
	"Popping",
	"Boogaloo",
	"Scarecrow",
	"Tutting",
	"Waving",
	"Glides",
	"Puppet",
	"Crazy Legs",
	"Snaking",
	"Strobing",
	"Ticking",
	"Dime Stop",
	"Saccing",
	"Air Posing",
	"Robot",
	"Digits",
	"Sleepy Style",
	"Toyman",
	"Banging",
	"Strutting",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", homePage)
	http.HandleFunc("/public/", publicResources) // css, js, img files
	http.HandleFunc("/web/", publicResources)
	http.HandleFunc("/styles", stylesFunc)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Printf("HTTP request for %v\n", string(path))

	first, second := generateStyles()

	data := map[string]string{
		"first":  first,
		"second": second,
	}

	page, err := ioutil.ReadFile("views/home.template")

	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(string(page))

	if err == nil {
		tmpl.Execute(w, data)
	}
}

func publicResources(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Printf("HTTP request for %v\n", string(path))

	page, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(string(path), ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(page)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error 404 - " + http.StatusText(404)))
	}
}

func stylesFunc(w http.ResponseWriter, req *http.Request) {
	first, second := generateStyles()
	styleMap := map[string]string{
		"first":  first,
		"second": second,
	}

	ret, _ := json.Marshal(styleMap)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(ret))
}

func generateStyles() (string, string) {
	first := rand.Intn(len(styles))
	second := rand.Intn(len(styles))

	// first and second should not be the same style
	for first == second {
		second = rand.Intn(len(styles))
	}

	return styles[first], styles[second]
}
