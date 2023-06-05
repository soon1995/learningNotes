package main

import "fmt"

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
