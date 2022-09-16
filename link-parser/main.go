package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("./testfiles/ex4.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// DFS recursive call on DOM Tree
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var link Link
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link.Href = attr.Val
				}
			}
			link.Text = n.FirstChild.Data
			fmt.Println(link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

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
