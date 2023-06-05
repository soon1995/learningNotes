// Use sync.Mutex to make bzip2.writer safe for concurrent
// use by multiple goroutines.
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	w := NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
  if err := w.Close(); err != nil {
    log.Fatalf("bzipper close: %v\n", err)
  }
}

// go build 
