package parser

import (
	"log"
	"strings"

	"code.google.com/p/go.net/html"
)

func GetLinksFrom(body string) []string {
	links := []string{}
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}
