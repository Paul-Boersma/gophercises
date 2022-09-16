package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/Paul-Boersma/gophercises/cyoa/model"
)

func main() {
	// Command line input
	filename := flag.String("file", "./stories/gopher.json", "Choose a .json story")
	flag.Parse()
	storyname := path.Base(*filename)[:1+len(path.Ext(*filename))]

	fmt.Printf("Welcome to story: %s\n\n", storyname)

	// Json File
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Go Struct
	story, err := model.JsonStory(file)
	if err != nil {
		log.Fatal(err)
	}

	// WebTemplate
	fmt.Println(story)
	tmpl := template.Must(template.ParseFiles("./templates/cyoa.gohtml"))
	tmpl.Execute(os.Stdout, story)

	// Parse the byteslice throught json.Unmarshal()
	// Create a template file
	// Use the go datastructures as input data for the templates
	// Set up a simple server and client
	// When a user makes a web request send the template as a response.

	// Setup webserver
	handlerCyoa := NewHandler(tmpl, story)
	http.HandleFunc("/cyoa", handlerCyoa.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func NewHandler(tmpl *template.Template, story model.Story) CyoaHTTPHandler {
	return CyoaHTTPHandler{
		story: story,
		tmpl:  tmpl,
	}
}

type CyoaHTTPHandler struct {
	story model.Story
	tmpl  *template.Template
}

func (h *CyoaHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.tmpl.Execute(w, h.story["intro"])
}
