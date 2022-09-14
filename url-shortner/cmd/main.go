package main

import (
	"fmt"
	"net/http"

	urlShort "github.com/Paul-Boersma/gophercises/url-shortner/pkg"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlShort.MapHandler(pathsToUrls, mux)
	fmt.Println("Starting the server on :8080")

	// The mapHandler is a Handler interface which functions as a ServeMux
	http.ListenAndServe(":8080", mapHandler)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// yaml := `
	// - path: /urlshort
	// url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	// url: https://github.com/gophercises/urlshort/tree/solution
	// `
	// yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
