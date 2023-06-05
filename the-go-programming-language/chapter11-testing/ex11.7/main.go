package main

import "fmt"

func main() {
	m := make(M)
	m.Add(1)
	m.Add(3)
	m1 := make(M)
	m1.Add(1)
	m1.Add(2)
	m.UnionWIth(m1)
	fmt.Println(m)

}
