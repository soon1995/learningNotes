package main

import (
	"io"
	"log"
	"net"
	"os"
)

// gopl.io/ch8/netcat1
// Netcat1 is a read-only TCP client.
func Netcat1() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func Netcat2() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

// gopl.io/ch8/netcat3
func Netcat3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} //signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func main() {
	// Netcat1()
	// Netcat2()
	Netcat3()
}
