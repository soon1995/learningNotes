// The startElement and endElement functions in gopl.io/ch5/outline2
// share a global variable, depth. Turn them into anonymous functions that share a
// variable local to the outline function
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  if err != nil {
    log.Fatalf("parsing html failed: %v", err)
  }
	var depth int
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(doc, pre, post)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
