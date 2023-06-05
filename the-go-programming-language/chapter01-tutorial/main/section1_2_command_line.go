package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	Echo1()
	Echo2()
	Echo3()
	fmt.Println(os.Args[1:])
	Exercise1_1()
	Exercise1_2()
	Exercise1_3()
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

// Exercise 1.1: Modify the echo program to also print os.Args[0],
// the name of the command that invoked it
func Exercise1_1() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

// Exercise 1.2: Modify the echo program to print the index and value
// of each of its arguments, one per line
func Exercise1_2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join
func Exercise1_3() {
	// go run main/main.go a b c aaaaaa aaaaaa a a
	// a a a a a a a a a a a a a a a a a a a a a
	t1 := time.Now()
	Echo1()
	fmt.Println(time.Now().Sub(t1)) // 2.48µs
	t2 := time.Now()
	Echo2()
	fmt.Println(time.Now().Sub(t2)) // 5.04µs
	t3 := time.Now()
	Echo3()
	fmt.Println(time.Now().Sub(t3)) // 1.32µs
}
