// Write a variadic function ElementsByTagName that, given a HTML node tree
// and zero or more names, returns all the elements that match one of those names.
// Here are two example calls:
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
package main

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	m := make(map[string]bool)
	for _, v := range name {
		m[v] = true
	}
	var list []*html.Node
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if _, ok := m[n.Data]; ok {
				list = append(list, n)
			}
		}
	}
	forNodeVisit(doc, pre, nil)
	return list
}

func forNodeVisit(n *html.Node, pre func(*html.Node), post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forNodeVisit(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
