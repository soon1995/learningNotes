// Copied from https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-9-concurrency-with-shared-variables.html
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string, cancel <-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

type Func func(key string, cancel <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Duplicate suppresion
type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f      Func
	mu     sync.Mutex // guards cache
	cache  map[string]*entry
	delete chan string
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry), delete: make(chan string)}
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key, cancel)
		close(e.ready)
	} else {
		// this is a repead request for this key
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}

	select {
	case <-cancel:
		memo.mu.Lock()
		delete(memo.cache, key)
		memo.mu.Unlock()
		return nil, fmt.Errorf("%s is canceled", key)
	default:
	}
	return e.res.value, e.res.err
}

func main() {
	m := New(httpGetBody)
	cancel := make(chan struct{})
	var n sync.WaitGroup
	go func() {
		timer := time.NewTimer(600 * time.Millisecond)
		<-timer.C
		close(cancel)
		timer.Stop()
	}()
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func incomingURLs() []string {
	return []string{
		"https://google.com",
		"https://google.com",
		"https://google.com",
		"https://google.com",
		"https://golang.org",
		"https://golang.org",
	}
}
