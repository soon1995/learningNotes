// Define a variadic (*IntSet).AddAll(...int) method that allows a list of
// values to be added, such as s.AddAll(1, 2, 3)/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.AddAll(155, 156, 156)
	fmt.Println(x.String()) // "{1 9 42 144}"
}

type IntSet struct {
	words []uint64
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds alist of values
func (s *IntSet) AddAll(values ...int) {
	for _, x := range values {
		s.Add(x)
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
