// Modify the reverb2 server to use a sync.WaitGroup per connection to count
// the number of active echo goroutines. When it falls to zero, close the write
// half of the TCP connection as described in Exercise 8.3. Verify that your
// modified netcat3 client from that exercise waits for the final echoes of multiple concurrent shouts,
// even after the standard input has been closed.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
		// reverb 2: go
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func Echo1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func main() {
	Echo1()
}
