// Implement these addtional method
// func (*IntSet) Len() int // return the number of elements
// func (*IntSet) Remove(x int) // remove x from the set
// func (*IntSet) Clear() // remove all elements from set
// func (*IntSet) Copy() *IntSet // return a copy of the set
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x) // "{1 9 42 144}"
	fmt.Println(x)  // "{[4398046511618 0 65535]}"

	z := x.Copy()
	z.Add(666)
	fmt.Println(&x) // "{1 9 42 144}"
	fmt.Println(z)  // "{1 9 42 144 666}"
	fmt.Println(x.Len()) // 4
	fmt.Println(z.Len()) // 5

  z.Remove(9)
  z.Remove(666)
  z.Remove(666)
  fmt.Println(z) // "{1 42 144}"
	fmt.Println(*z) // {[4398046511106 0 65536 0 0 0 0 0 0 0 0]}

	z.Clear()
	fmt.Println(&x) // "{1 9 42 144}"
	fmt.Println(z)  // "{}"


}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count

}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
  s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	y := make([]uint64, len(s.words))
	copy(y, s.words)
	return &IntSet{y}
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
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
