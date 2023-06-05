// gopl.io/ch9/memo5
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

type entry struct {
	res   result
	ready chan struct{}
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, name string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(name)
	// Broadcast the ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition
	<-e.ready
	// Send the result to the client
	response <- e.res
}

func main() {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
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
