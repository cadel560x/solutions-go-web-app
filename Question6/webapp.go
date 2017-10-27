package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type goPlaceholders struct {
	Message   string
	UserGuess string
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerGame(w http.ResponseWriter, r *http.Request) {
	// Initializing local variables
	guess := r.FormValue("guess")
	placeholders := &goPlaceholders{Message: "Guess a number between 1 and 20"}
	rand.Seed(time.Now().UnixNano())
	randomNumber := strconv.Itoa(rand.Intn(20-1) + 1)

	// Setting up cookie 'target'
	if _, err := r.Cookie("target"); err != nil {
		http.SetCookie(w, &http.Cookie{Name: "target", Value: randomNumber})
	}

	// If 'guess' variable has something, display its value
	if len(guess) > 0 {
		placeholders.UserGuess = "You guessed " + guess
	}

	templ, _ := template.ParseFiles("guess.tmpl")
	templ.Execute(w, placeholders)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerGame)
	http.ListenAndServe(":8080", nil)
}
