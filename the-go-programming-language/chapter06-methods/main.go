package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"
	"net/url"
	"sync"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(i float64) {
	p.X *= i
	p.Y *= i
}

// A path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance travelled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// Values maps a string key to a list of values
type Values map[string][]string

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key
// It appends to any existing values associated with key/
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func URLValues() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"
	fmt.Println(m["item"])     // [1 2]
	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint1 struct {
	*Point
	Color color.RGBA
}

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping: make(map[string]string)}

func Lookup1(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
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

func main() {
	// perim := Path{{1,1}, {5,1}, {5,4},{1,1}}
	// fmt.Println(geometry.PathDistance(perim))
	// fmt.Println(perim.Distance())
	// p := Point{1, 2}
	// p.ScaleBy(2)
	// fmt.Println(p) // "{2, 4}"

	//  red := color.RGBA{255,0,0,255}
	//  blue := color.RGBA{0,0,255,255}
	// p := ColoredPoint1{&Point{1, 1}, red}
	// q := ColoredPoint1{&Point{5, 4}, blue}
	// q.Point = p.Point
	// p.ScaleBy(2)
	// fmt.Println(*p.Point, *q.Point)

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
  fmt.Println(x) // "{[4398046511618 0 65535]}"
}
