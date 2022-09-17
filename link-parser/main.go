package main

import (
	"github.com/Paul-Boersma/gophercises/pkg/parser"
)

func main() {
	parser.ParseHTML("./testfiles/ex1.html")
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
