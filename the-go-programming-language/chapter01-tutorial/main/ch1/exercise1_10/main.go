package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// Exercise 1.10: Find a web site that produces a large amount of data.
// Investigate caching by running fetchall twice in succession to see whether the reported
// time changes much.  Do  you get the same content each time? Modify fetchall to print its
// output to a file so it can be examined.
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		fmt.Printf("start fetch %v\n", url)
		go fetch(url, ch)
		go fetch(url, ch)
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

	filepath := uuid.NewString()
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ch <- fmt.Sprintf("while open %s: %v", filepath, err)
		return
	}
  defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
