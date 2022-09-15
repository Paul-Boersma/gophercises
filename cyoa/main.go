package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Paul-Boersma/gophercises/cyoa/model"
)

func main() {
	filename := flag.String("file", "./stories/gopher.json", "Choose a .json story")
	flag.Parse()
	storyname := path.Base(*filename)[:1+len(path.Ext(*filename))]

	fmt.Printf("Welcome to story: %s\n", storyname)

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var story model.Story
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&story); err != nil {
		log.Fatal(err)
	}

	fmt.Println(story)

	// Parse the byteslice throught json.Unmarshal()
	// Create a template file
	// Use the go datastructures as input data for the templates
	// Set up a simple server and client
	// When a user makes a web request send the template as a response.

}
