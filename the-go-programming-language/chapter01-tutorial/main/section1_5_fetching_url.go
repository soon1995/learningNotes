package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Fetch()
	// Exercise1_7()
	// Exercise1_8()
	// Exercise1_9()
	FetchAll()
}

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}

// The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout
// without requiring a buffer large enough to hold the entire stream. Be sure
// to check the error result of io.Copy.
func Exercise1_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.
func Exercise1_8() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}

// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
func Exercise1_9() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s -- [%s]", b, resp.Status)
	}
}

// fetch many URLs
// discards the responses but reports the size and elapsed time for each one
// Fetchall fetches URLs in parallel and reports their times and sizes.
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		fmt.Printf("start fetch %v\n", url)
		// go fetch(url, ch)
		// go fetch(url, ch)
		go Exercise1_10(url, ch)
		go Exercise1_10(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// Exercise 1.10: Find a web site that produces a large amount of data.
// Investigate caching by running fetchall twice in succession to see whether the reported
// time changes much.  Do  you get the same content each time? Modify fetchall to print its
// output to a file so it can be examined.
func Exercise1_10(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	filepath := uuid.NewString()
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ch <- fmt.Sprintf("while while open %s: %v", &filepath, err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
