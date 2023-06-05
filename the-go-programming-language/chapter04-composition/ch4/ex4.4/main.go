// Write a version of rotate that operates in a single pass
package main

import "fmt"

func rotate(s []int) {
	temp := s[0]
	copy(s, s[1:])
	s[len(s)-1] = temp
}

func rotateEx() {
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
