// Implement countWordsAndImages (See Ex4.9 for word splitting)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: count URL\n")
		os.Exit(1)
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%d words, %d images", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words2, images2 := countWordsAndImages(c)
		words += words2
		images += images2
	}
	switch n.Type {
	case html.TextNode:
		words += countWords(n.Data)
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	}
	return
}

func countWords(s string) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return
}
