package parser

import (
	"os"
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

	return parseAnchorTags(doc)
}

func parseAnchorTags(n *html.Node) (links []Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var link Link
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val
			}
		}
		link.Text = n.FirstChild.Data
		links = append(links, link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseAnchorTags(c)
	}
}
