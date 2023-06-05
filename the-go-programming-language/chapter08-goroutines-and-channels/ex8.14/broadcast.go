package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

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

func handleConnChat(conn net.Conn, who string) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

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
