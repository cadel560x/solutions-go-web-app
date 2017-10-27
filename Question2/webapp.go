package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<!doctype html>\n"+
		"<html lang=\"en\">\n"+
		" <head>\n"+
		"   <meta charset=\"utf-8\">\n"+
		"   <title>Guessing Game</title>\n"+
		" </head>\n"+
		" <body>\n"+
		"   <h1>Guessing Game</h1>\n"+
		" </body>\n"+
		"</html>\n")

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
