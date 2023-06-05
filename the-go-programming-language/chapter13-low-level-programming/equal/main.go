package main

import "fmt"

func main() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
	fmt.Println(Equal([]string(nil), []string{}))
	// fmt.Println(Equal(map[string]int(nil), map[string]int{}))

	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a))
	fmt.Println(Equal(b, b))
	fmt.Println(Equal(c, c))
	fmt.Println(Equal(a, b))
	fmt.Println(Equal(a, c))

}

type link struct {
	value string
	tail  *link
}
