package main

import (
	"fmt"
	"net"
	"time"

	"alsoon.com/go_code/chatsystem/server/model"
)

// go routine for each connection
func connProcess(conn net.Conn) {
	defer conn.Close()

	p := *&Processor{
		Conn: conn,
	}

	err := p.processing()
	if err != nil {
		// fmt.Println("Err connection: ", err)
		return
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	initPool("x.x.x.x:6379", 8, 0, 300*time.Second)
	initUserDao()

	fmt.Println("Server port 8888 is listening...")
	l, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Failed starting - err:", err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to make connection err=", err)
		}
		// if connection was made, start a routine for each connection
		go connProcess(conn)
	}
}
