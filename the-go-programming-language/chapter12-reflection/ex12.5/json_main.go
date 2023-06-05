package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Wrapper struct {
	item Item
	c    complex128
}

type Item struct {
	v int
}

func main() {
	w := &Wrapper{c: complex(1, 2)}
	b, err := json.MarshalIndent(w, "", " ")
	if err != nil {
		log.Fatalf("failed marshal indent")
	}
	fmt.Println(string(b))
}
