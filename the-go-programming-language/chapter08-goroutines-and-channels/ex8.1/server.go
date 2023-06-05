// Modify clock2 to accept a port number, and write a program, clockwall,
// that acts as a client of several clock servers at once, reading the times
// from each one and displaying the results in a table, akin to the wall of clocks
// seen in some business offices. If you have access to geopraphically distributed computers,
// run instances remotely; otherwise run local instances on different ports
// with fake time zones.
// $ TZ=US/Eastern ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8000 London=localhost:8010 Tokyo=localhost:8020
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func Clock2() {
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnecter
		}
		time.Sleep(1 * time.Second)
	}
}

var (
	port     = flag.String("port", "", "port")
	timezone = flag.String("tz", "", "timezone e.g. US/Eastern")
)

func main() {
	flag.Parse()
	if *port == "" || *timezone == "" {
		log.Fatal("Usage: cmd -port 8000 -tz US/Eastern")
	}
	os.Setenv("TZ", *timezone)
	Clock2()
}
