package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join
func main() {
	// go run main/main.go asdfsadf asdfasdf asdfasdf asdfasdf adsf asdf adsf adf asdf
	t1 := time.Now()
	Echo1()
	fmt.Println(time.Now().Sub(t1)) // 28.23µs
	t2 := time.Now()
	Echo2()
	fmt.Println(time.Now().Sub(t2)) // 3.78µs
	t3 := time.Now()
	Echo3()
	fmt.Println(time.Now().Sub(t3)) // 1.14µs
}

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// it has the best performance, it use []byte
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
