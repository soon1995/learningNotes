package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, host string
}

func (c *clock) Netcat1() {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c.watch(os.Stdout, conn)
}

func (c *clock) watch(dst io.Writer, src io.Reader) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, "%s\t: %s\n", c.name, s.Text())
	}
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

var m map[string]string

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(0)
	}
	m = make(map[string]string)
	for _, arg := range args {
		s := strings.Split(arg, "=")
		if len(s) != 2 {
			log.Fatalf("unexpected args %s", arg)
		}
		m[s[0]] = s[1]
	}
	for zone, host := range m {
		c := &clock{zone, host}
		go c.Netcat1()
	}

	time.Sleep(100 * time.Second)
}
