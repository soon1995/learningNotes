// Write a function to populate a mapping from element names --p, div, span and
// so on-- to the number of elements with that name in an HTML document tree
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// counts, err := CountTag(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "count tag: %v\n", err)
	// 	os.Exit(1)
	// }
	// for k, v := range counts {
	// 	fmt.Fprintf(os.Stdout, "Tag %-10s\t%d\n", k, v)
	// }

	main1()
}

func CountTag(r io.Reader) (map[string]int, error) {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int)
	countTag(counts, doc)
	return counts, nil
}

func countTag(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTag(counts, c)
	}
}

func main1() {
	freq, err := tagFreq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}

func tagFreq(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)
	z := html.NewTokenizer(os.Stdin)
	var err error
	for {
		type_ := z.Next()
		if type_ == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}
