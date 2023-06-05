// In netcat3, the interface value conn has the concrete type *net.TCPConn, which
// represents a TCP connection. A TCP connection consists of two halves that may be closed
// independently using its CloseRead and CloseWrite methods. Modify the main goroutine of
// netcat to close only the write half of the connection so that the program will continue
// to print the final echoes from the reverb1 server even after the standard input has been closed.
// (Doing ths for the reverb2 server is harder; see Exercise 8.4)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
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
		go handleConn2(conn)
	}
}

func main() {
	Echo1()
}
