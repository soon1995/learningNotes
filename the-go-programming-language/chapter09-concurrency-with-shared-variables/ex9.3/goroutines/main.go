// Extend the Func type and the (*Memo).Get method so that callers may provide an optional done channel through which
// they can cancel the operation. The results of a cancelled Func call should not be cached.
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

type request struct {
	key      string
	response chan<- result
	cancel   <-chan struct{}
}

type Memo struct {
	requests chan request
	delete   chan string
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), delete: make(chan string)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Delete(key string) {
	memo.delete <- key
}

func (memo *Memo) Close() {
	close(memo.requests)
	close(memo.delete)
}

type entry struct {
	res   result
	ready chan struct{}
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
loop:
	for {
		select {
		case key, ok := <-memo.delete:
			if !ok {
				break loop
			}
			delete(cache, key)
		case req, ok := <-memo.requests:
			if !ok {
				break loop
			}
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.cancel)
			}
			go e.deliver(req.response, req.cancel, memo.delete, req.key)
		}
	}
}

func (e *entry) call(f Func, name string, cancel <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(name, cancel)
	// Broadcast the ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, cancel <-chan struct{}, delete chan<- string, key string) {
	// Wait for the ready condition
	<-e.ready

	select {
	case <-cancel:
		delete <- key
		response <- result{nil, fmt.Errorf("%s is canceled", key)}
	default:
		// Send the result to the client.
		response <- e.res
	}
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
