# Exploring `slice`

Q: Does modifying the underlying elements in a slice of function arguments affect the original slice?

```go
func changeElem(s []int, n int) {
	for i := range s {
		s[i] = n
	}
}

func main() {
	s := make([]int, 5)
	changeElem(s, 2)
	fmt.Println(s) // [2 2 2 2 2]
}
```

Yes, modifying the underlying elements in a slice of function arguments will affect the original slice.

Q: How about appending?

```go
func appendElem(s []int, n int) {
	s = append(s, n)
}

func appendElemByPointer(s *[]int, n int) {
	*s = append(*s, n)
}

func main() {
	s := []int{2,2,2,2,2}
	appendElem(s, 3)
	fmt.Println(s) // [2 2 2 2 2]

	appendElemByPointer(&s, 3)
	fmt.Println(s) // [2 2 2 2 2 3]
}
```

Yes if the argument receives pointer slice; No if it does not.

Q: The benchmarks of finding non-empty elements in a slice

```golang
// in-place slice algorithm, it also changes the slice that passed in
func nonemptyInplace1(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// also in-place slice algorithm, it also changes the slice that passed in
func nonemptyInplace2(strings []string) []string {
	out := strings[:0] // zero-length slice of original, but cap is the cap(strings)
	for _, s := range strings {
		if s != "" {
			out = append(out, s) // thus appending does not append the len
		}
	}
	return out
}

func nonemptyWithNewSlice(strings []string) []string {
	out := []string{}
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
```

Result:

```bash
goos: linux
goarch: amd64
pkg: testGO
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkNonEmptyInplace1-4             291026764                4.144 ns/op
BenchmarkNonEmptyInplace2-4             224409214                5.271 ns/op
BenchmarkNonEmptyWithNewSlice-4         27945484                42.34 ns/op
PASS
```
