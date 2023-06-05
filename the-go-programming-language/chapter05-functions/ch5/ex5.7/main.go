// Develop startElement and endElement into a general HTML pretty-printer
// Print comment nodes, text nodes, and the attributes of each element(<a href='...'>).
// Use short forms like <img/> instead of <img></img> when an element has no children.
// Write a test to ensure that the output can be parsed successfully. (See Chapter 11)
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
  p := NewPrettier(os.Stdout)
  p.Pretty(os.Stdin)
}

type Prettier struct {
	io.Writer
}

func NewPrettier(out io.Writer) *Prettier {
  return &Prettier{out}
}

func (p *Prettier) Pretty(in io.Reader) {
	doc, err := html.Parse(in)
	if err != nil {
		log.Fatalf("parse failed: %v", err)
	}
	p.forEachNode(doc, p.startElement, p.endElement)
}

func (p *Prettier) writef(format string, a ...interface{}) {
	p.Write([]byte(fmt.Sprintf(format, a...)))
}

func (p *Prettier) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func (p *Prettier) startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		p.startElementNode(n)
	case html.TextNode:
		s := strings.TrimSpace(n.Data)
		if s != "" {
			p.writef("%*s%s\n", depth*2, "", s)
		}
	case html.CommentNode:
		s := strings.TrimSpace(n.Data)
		if s != "" {
			p.writef("%*s%s\n", depth*2, "", s)
		}
	}
}

func (p *Prettier) startElementNode(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	attributes := bytes.NewBuffer(nil)
	for i, attr := range n.Attr {
		if i == 0 {
			attributes.WriteString(" ")
		}
		attributes.WriteString(attr.Key + "='" + attr.Val + "'")
		if i < len(n.Attr)-1 {
			attributes.WriteString(" ")
		}
	}
	p.writef("%*s<%s%s%s\n", depth*2, "", n.Data, attributes.String(), end)
	depth++
}

func (p *Prettier) endElement(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	depth--
	if n.FirstChild == nil {
		return
	}
	p.writef("%*s</%s>\n", depth*2, "", n.Data)
}
