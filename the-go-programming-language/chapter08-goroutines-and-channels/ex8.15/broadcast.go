// Failure of any client program to read data in a timely manner ultimately causes all clients
// to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not
// ready to accept it. Alternatively, add buffering to each client's outgoing message channel so that
// most messages are not dropped; the broadcaster should use a non-blocking send to this channel.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string // an outgoing message channel

const bufSize = 10

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client message
)

func broadcaster() {
	clients := make(map[client]bool) // all connected client
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels
			for cli := range clients {
				fmt.Println(len(cli))
				if len(cli) < bufSize {
					cli <- msg
				}
				// select {
				// case cli <- msg:
				// default:
				// 	// Skip client if it's reading messages slowly.
				// }
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConnChat(conn net.Conn) {
	ch := make(chan string, bufSize) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	maxIdle := 1 * time.Minute
	timeout := time.NewTicker(maxIdle)
	go func() {
		<-timeout.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(maxIdle)
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func Chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnChat(conn)
	}
}
