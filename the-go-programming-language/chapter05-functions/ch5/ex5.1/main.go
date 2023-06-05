// Change the findlinks program to traverse the n.FirstChild linked list using
// recursive calls to visit instead of loop
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
