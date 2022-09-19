package main

import (
	"encoding/xml"
	"flag"
	"github.com/Paul-Boersma/gophercises"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}

// No cyclical dependencies
// Provide map in XML format
// Use link-parser package
