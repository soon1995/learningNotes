// Modify crawl to make local copies of the pages it finds, creating directories
// as necessary. Don't make copies of pages that come from a different domain.
// For example, if the original page comes from golang.org, save all files from
// there, but exclude ones from vimeo.com
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ex5.13 URL..")
	}
	for _, arg := range os.Args[1:] {
		links := crawl(arg)
		for _, link := range links {
			err := saveFile(link)
			if err != nil {
				fmt.Printf("cannot cache %s: %v", link, err)
			}
		}
	}

	// https://golang.org
	// breadthFirst(crawl, os.Args[1:])
}

func saveFile(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, url.Path)
		filename = url.Path
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// package links provides a link-extraction function
// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	exist := make(map[string]bool)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				if exist[a.Val] {
					continue
				}
				exist[a.Val] = true
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URL
				}
				parent, _ := resp.Request.URL.Parse(url)
				if link.Hostname() == parent.Hostname() {
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
