// Modify forEachNode so that the pre and post functions return a boolean result
// indicating whether to continue the traversal. Use it to write a function
// ElementByID with the following signature that finds the first HTML element with
// the specified id attribute. The function should stop the traversal as soon as a
// match is found.
// func ElementByID(doc *html.Node, id string) *html.Node
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, v := range n.Attr {
			if v.Key == "id" && v.Val == id {
				return true
			}
		}
		return false
	}
	return forEachNode(doc, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		if post(n) {
			return n
		}
	}
	return nil
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("parse failed", err)
	}
	if len(os.Args) != 2 {
		log.Fatal("usage: program ID")
	}
	id := os.Args[1]
	node := ElementByID(doc, id)
  fmt.Printf("%+v\n", node)
  fmt.Printf("%#v", node)
	if node == nil {
		fmt.Println("no element has id", id)
	} else {
		fmt.Printf("element <%s> has id %s\n", node.Data, id)
	}
}
