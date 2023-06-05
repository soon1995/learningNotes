// Write a program with two goroutines that send messages back and forth
// over two unbuffered channels in ping-pong fashion. How many communications
// per second can the program sustain?
// 1 min = 116,885,834: 1,948,097 communications per second
package main

import (
	"fmt"
	"time"
)

var (
	ch1    = make(chan string)
	ch2    = make(chan string)
	count  int64
	cancel = make(chan struct{})
)

func main() {
	go func() {
		for {
			select {
			case <-cancel:
				return
			case res := <-ch1:
				ch2 <- res
			}
		}
	}()
	go func() {
		for {
			select {
			case <-cancel:
				return
			case res := <-ch2:
				count++
				ch1 <- res
			}
		}
	}()
	go func() {
		timer := time.NewTimer(1 * time.Minute)
		<-timer.C
		close(cancel)
		timer.Stop()
	}()
	ch1 <- "hi"

loop:
	for {
		select {
		case <-cancel:
			break loop
		}
	}
	select {
	case <-ch1:
	case <-ch2:
	}
	fmt.Println(count)
}
