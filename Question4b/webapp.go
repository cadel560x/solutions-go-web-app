package main

import (
	"html/template"
	"net/http"
)

type messageT struct {
	Message string
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerGame(w http.ResponseWriter, r *http.Request) {
	message := &messageT{Message: "Guess a number between 1 and 20"}

	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, message)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerGame)
	http.ListenAndServe(":8080", nil)
}
