package main

import (
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
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

	output := styles[first] + " + " + styles[second]
	w.Write([]byte(output))
}
