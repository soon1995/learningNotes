// refer torbiak/gopl.io/ex8.10
// HTTP requests may be cancelled by closing the optional Cancel channel in the
// http.Request struct. Modify the web crawler of Section 8.6 to support cancellation.
// Hint: the http.Get convenience function does not give you an opportunity to customize a Request.
// Instead, create the request using http.NewRequest, set its Cancel field, then perform the request
// by calling http.DefaultClient.Do(req)
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

var maxDepth = flag.Int("depth", -1, "max crawl depth")
var tokens = make(chan struct{}, 20)
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

var cancel = make(chan struct{})

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= *maxDepth {
		return
	}
	tokens <- struct{}{}
	list, err := Extract(url, cancel)
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
	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})
	go func() {
		signal.Notify(interrupt, os.Interrupt)
	}()
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return
	case <-interrupt:
		close(cancel)
	}
}
