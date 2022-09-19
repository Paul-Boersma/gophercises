package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	parser "github.com/Paul-Boersma/gophercises/link-parser/pkg"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "https://calhoun.io", "url for which a sitemap is built")
	flag.Parse()

	assureWebProtocol(&url)

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parser.ParseHtmlReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	for i, link := range links {
		fmt.Printf("%v: %v --> %v\n", i, link.Href, link.Text)
	}
}

func bfsSiteBuilder(url string, memo map[string][]string) map[string][]string {
	if url == "" {
		return nil
	}
	if _, ok := memo[url]; ok {
		return memo
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	links, err := parser.ParseHtmlReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		return bfsSiteBuilder(link.Href, memo)
	}

	return memo
}

func assureWebProtocol(url *string) {
	hasWebProtocol := strings.HasPrefix(*url, "https://") || strings.HasPrefix(*url, "http://")
	if !hasWebProtocol {
		*url = "https://" + *url
	}
}
