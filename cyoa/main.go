package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Paul-Boersma/gophercises/cyoa/model"
)

func main() {
	filename := flag.String("file", "./stories/gopher.json", "Choose a .json story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	story, err := model.JsonStory(file)
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("."))
	storyHandler := NewHandler(story)

	http.HandleFunc("/", storyHandler.ServeHTTP)
	http.Handle("/templates/", http.StripPrefix("", fs))

	http.ListenAndServe(":8080", nil)
}
