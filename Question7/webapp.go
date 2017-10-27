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
	Message string
	// As seen in https://stackoverflow.com/questions/18175630/go-template-executetemplate-include-html
	UserGuess template.HTML
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
	outputHTML := ""

	// Check if 'guess' variable and 'target' cookie are set
	if len(guess) > 0 && len(targetCookie.Value) > 0 {
		// User wins
		if strings.Compare(targetCookie.Value, guess) == 0 {
			outputHTML += "<h3>Congratulations!! You've guessed the number!!</h3>\n"
			outputHTML += "<h3>The number was " + guess + "</h3>\n"
			outputHTML += "<h3>Want to try again? <a href=\"/guess\">New Game</a></3>\n"

			// Set new cookie for the next game
			http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(20-1) + 1)})

			// User loses
		} else {
			outputHTML += "<h3>You guessed " + guess + "</h3>\n"

			// Guess is high
			if guess > targetCookie.Value {
				outputHTML += "<h3>Your guess is high</h3>\n"
				// Guess is low
			} else {
				outputHTML += "<h3>Your guess is low</h3>\n"
			}
		} // if - else

		placeholders.UserGuess = template.HTML(outputHTML)
		// Set cookie if not set
	} else if targetCookie == nil {
		http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(20-1) + 1)})
	}

	templ, _ := template.ParseFiles("guess.tmpl")
	templ.Execute(w, placeholders)

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/guess", handlerGame)
	http.ListenAndServe(":8080", nil)
}
