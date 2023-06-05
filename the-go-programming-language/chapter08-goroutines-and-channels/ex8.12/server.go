// Make the broadcaster announce the current set of clients to each new arrival.
// This requires that the clients set and the entering and leaving channels
// record the client name too.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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

type client struct {
	outgoing chan<- string // an outgoing message channel
	name     string
}

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
				cli.outgoing <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.outgoing <- "online users:"
			for c := range clients {
				if cli != c {
					cli.outgoing <- c.name
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outgoing)
		}
	}
}

func handleConnChat(conn net.Conn) {
	var cli client
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	cli.outgoing = ch
	cli.name = who

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	Chat()
}
