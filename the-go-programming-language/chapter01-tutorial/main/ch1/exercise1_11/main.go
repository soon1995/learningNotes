package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Exercise 1.11: Try fetchall with longer argument lists, such as samples from the top
// million web sites available at alexa.com. How does the program behave if a web site
// just doesn’t respond?  (Section 8.9 describes mechanisms for coping in such cases.)
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		fmt.Printf("start fetch %v\n", url)
		go fetch(url, ch)
	}
	for range os.Args[1:] {
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

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s %s", secs, nbytes, url, resp.Status)
}
