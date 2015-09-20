package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	first, second := generateStyles()

	data := map[string]string{
		"first":  first,
		"second": second,
	}

	page, err := ioutil.ReadFile("templates/home.template")

	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(string(page))
	if err == nil {
		tmpl.Execute(w, data)
	}
}

func generateStyles() (string, string) {
	styles := []string{
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
	}

	first := rand.Intn(len(styles))
	second := rand.Intn(len(styles))

	// first and second should not be the same style
	for first == second {
		second = rand.Intn(len(styles))
	}

	return styles[first], styles[second]
}
