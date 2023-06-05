// Using a select statement, add a timeout to the echo server from Section 8.3
// so that it disconnects any client that shouts nothing within 10 seconds
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

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	ticker := time.NewTicker(10 * time.Second)
	text := make(chan string)
	var wg sync.WaitGroup
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		close(text)
	}()
	for {
		select {
		case t, ok := <-text:
			if ok {
				wg.Add(1)
				ticker.Reset(10 * time.Second)
				go func() {
					defer wg.Done()
					echo(c, t, 1*time.Second)
				}()
			} else {
				wg.Wait()
				c.Close()
				return
			}
		case <-ticker.C:
			ticker.Stop()
			c.Close()
			fmt.Println("disconnect silent client")
			return
		}
	}
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
		go handleConn2(conn)
	}
}

func main() {
	Echo1()
}
