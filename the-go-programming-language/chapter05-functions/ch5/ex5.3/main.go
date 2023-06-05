// Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into <script> or <style> element, since their contents are not
// visible in a web browser
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Printext %s", err)
		os.Exit(1)
	}
	printText(doc)
  // main1()
}

func printText(n *html.Node) {
	if n.Type == html.TextNode {
		text := n.Data
		if len(strings.TrimSpace(text)) > 0 {
			fmt.Print(text)
      if text[len(text) - 1] != '\n' {
        fmt.Println()
      }
		}
	}
	if n.Type == html.ElementNode {
		tag := n.Data
		if tag == "script" || tag == "style" {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c)
	}
}

func printTagText(r io.Reader, w io.Writer) error {
	z := html.NewTokenizer(os.Stdin)
	var err error
	stack := make([]string, 20)
Tokenize:
	for {
		switch z.Next() {
		case html.ErrorToken:
			break Tokenize
		case html.StartTagToken:
			b, _ := z.TagName()
			stack = append(stack, string(b))
		case html.TextToken:
			cur := stack[len(stack)-1]
			if cur == "script" || cur == "style" {
				continue
			}
			text := z.Text()
			if len(strings.TrimSpace(string(text))) == 0 {
				continue
			}
			w.Write([]byte(fmt.Sprintf("<%s>", cur)))
			w.Write(text)
			if text[len(text)-1] != '\n' {
				io.WriteString(w, "\n")
			}
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func main1() {
	err := printTagText(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
