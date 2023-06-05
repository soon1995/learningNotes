// Following the approach of mirroredQuery in Section 8.4.4, implement a variant
// of fetch that requests several URLs concurrently. As soon as the first response
// arrives, cancel the other request.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string, cancel <-chan struct{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func fetchFirst(urls []string) string {
	if len(urls) == 0 {
		return ""
	}
	responses := make(chan string, len(urls))
	cancel := make(chan struct{})
	for _, url := range urls {
		u := url
		go func() {
			err := fetch(u, cancel)
			if err == nil {
				responses <- u
			}
		}()
	}
	res := <-responses
	close(cancel)
	return res
}

func main() {
	fmt.Println("first fetch:", fetchFirst(os.Args[1:]))
}
