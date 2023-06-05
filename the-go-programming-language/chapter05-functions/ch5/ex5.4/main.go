// Extend the visit function so that it extracts other kinds of links from the document
// such as images, scripts, and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	Findlinks()
}

func Findlinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

// visit appends to links each link found in n and returns the result
func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				fmt.Printf("<%s> %s\n", n.Data,a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
