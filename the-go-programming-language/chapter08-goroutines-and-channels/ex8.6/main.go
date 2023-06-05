// Add depth-limiting to the concurrent crawler. That is, if the user sets -depth=3,
// then only URLs erachable by at most three links will be fetched.
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

var maxDepth = flag.Int("depth", -1, "max crawl depth")
var tokens = make(chan struct{}, 20)
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= *maxDepth {
		return
	}
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock()
		if !seen[link] {
			seen[link] = true
			seenLock.Unlock()
			wg.Add(1)
			go crawl(link, depth+1, wg)
		} else {
			seenLock.Unlock()
		}
	}
}

func main() {
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}
