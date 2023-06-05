// Make the chat server disconnect idle clients, such as those have sent
// no messages in the last five minutes. Hint: calling conn.Close() in
// oanother goroutine unblocks active Read calls such as the one done by input.Scan()
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	Chat()
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

type client chan<- string // an outgoing message channel

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
				cli <- msg
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
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	maxIdle := 5 * time.Second
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
