package main

import (
	"fmt"
	"log"

	parser "github.com/Paul-Boersma/gophercises/link-parser/pkg"
)

func main() {
	links, err := parser.ParseHTML("./testfiles/ex2.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(links)
}

// Wat wil ik testen?
// - Bevat het link struct de te verwachten href en text
// -

// html package makes use of a:
// - Tokenizer
// - Parser

// token: 'A SYMBOL RESPRESENTING SOMETHING'
// eg. A token is created when a user is logged in. This token represents the fact
// that the user has been logged in succesfully, so that the authentication process
// does not need to be reconsidered.
