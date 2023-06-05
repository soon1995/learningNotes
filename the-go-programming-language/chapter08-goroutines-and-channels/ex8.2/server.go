// Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret
// commands from each client such as cd to change directory, ls to list a directory, get
// to send the contents of a file, and close to close the connection. You can use the standard
// ftp command as the client, or write your own.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type FTPServer struct {
}

func (f *FTPServer) cd(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}

func (f *FTPServer) list(c net.Conn, path ...string) error {
	e := exec.Command("ls", path...)
	e.Stdout = c
	err := e.Run()
	if err != nil {
		return fmt.Errorf("bad path: %s", path)
	}
	return nil
}

func (f *FTPServer) get(c net.Conn, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	mustCopy(c, file)
	return nil
}

func (f *FTPServer) close(c net.Conn) error {
	c.Close()
	return nil
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	server := new(FTPServer)
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(c, server)
	}
}

func handleConn(c net.Conn, ftpServer *FTPServer) {
	defer c.Close()
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "ls":
			ftpServer.list(c, fields[1:]...)
		case "cd":
			ftpServer.cd(fields[1])
		case "get":
			ftpServer.get(c, fields[1])
		case "close":
			return
		default:
			fmt.Fprint(c, "usage:\nls [path...]\ncd path\nget filename\nclose\n\n")
		}
	}
}
