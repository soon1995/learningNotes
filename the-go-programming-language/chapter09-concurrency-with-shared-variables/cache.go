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

// gopl.io/ch9/mem1
// A Memo caches the results of a calling func
// type Memo struct {
// 	f     Func
// 	cache map[string]result
// }

// gopl.io/ch9/mem2
// type Memo struct {
// 	f     Func
// 	mu    sync.Mutex // guards cache
// 	cache map[string]result
// }

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// func New(f Func) *Memo {
// 	return &Memo{f: f, cache: make(map[string]result)}
// }

// NOTE: not concurrency-safe!
// gopl.io/ch9/mem1
// func (memo *Memo) Get(key string) (interface{}, error) {
// 	res, ok := memo.cache[key]
// 	if !ok {
// 		res.value, res.err = memo.f(key)
// 		memo.cache[key] = res
// 	}
// 	return res.value, res.err
// }

// gopl.io/ch9/mem2
// func (memo *Memo) Get(key string) (interface{}, error) {
// 	memo.mu.Lock()
// 	defer memo.mu.Unlock()
// 	res, ok := memo.cache[key]
// 	if !ok {
// 		res.value, res.err = memo.f(key)
// 		memo.cache[key] = res
// 	}
// 	return res.value, res.err
// }

// gopl.io/ch9/mem3
// after mem3, the performance improves again, but now we notice that some URLs are
// being fetched twice. This happens when two or more goroutines call Get for the same URL at
// about the same time.jjjjjjjj
// func (memo *Memo) Get(key string) (interface{}, error) {
// 	memo.mu.Lock()
// 	res, ok := memo.cache[key]
// 	memo.mu.Unlock()
// 	if !ok {
// 		res.value, res.err = memo.f(key)

// 		// Between the two critical sections, several goroutines
// 		// may race to comput f(key) and update the map
// 		memo.mu.Lock()
// 		memo.cache[key] = res
// 		memo.mu.Unlock()
// 	}
// 	return res.value, res.err
// }

// gopl.io/ch9/mem4
// Duplicate suppresion
type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		// this is a repead request for this key
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
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
