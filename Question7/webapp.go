package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type goPlaceholders struct {
	Message   string
	UserGuess string
	NewGame   string
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerGame(w http.ResponseWriter, r *http.Request) {
	// Initializing local variables
	targetCookie, _ := r.Cookie("target")
	guess := r.FormValue("guess")
	placeholders := &goPlaceholders{Message: "Guess a number between 1 and 20"}
	rand.Seed(time.Now().UnixNano())
	randomNumber := strconv.Itoa(rand.Intn(20-1) + 1)

	// Check if 'guess' variable and 'target' cookie are set
	if len(guess) > 0 && len(targetCookie.Value) > 0 {
		// User wins
		if strings.Compare(targetCookie.Value, guess) == 0 {
			placeholders.UserGuess = "Congratulations!! You've guessed the number " + guess + "!!!"
			placeholders.NewGame = "Want to try again? New Game"
			// User loses
		} else {
			placeholders.UserGuess = "You guessed " + guess + ". "
			if guess > targetCookie.Value {
				placeholders.UserGuess += "Your guess is high."
			} else {
				placeholders.UserGuess += "Your guess is low."
			}
		}
		// Set cookie if not set
	} else if targetCookie == nil {
		http.SetCookie(w, &http.Cookie{Name: "target", Value: randomNumber})
	}

	templ, _ := template.ParseFiles("guess.tmpl")
	templ.Execute(w, placeholders)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerGame)
	http.ListenAndServe(":8080", nil)
}
