package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHTML(filename string) (links []Link, err error) {
	if filepath.Ext(filename) != ".html" {
		fmt.Errorf("Provided file should be of format '.html'")
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	return parseAnchorTags(doc), nil
}

func parseAnchorTags(n *html.Node) (links []Link) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, Link{Href: attr.Val, Text: n.FirstChild.Data})
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseAnchorTags(c)...)
	}

	return links
}
