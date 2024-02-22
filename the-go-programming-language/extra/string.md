# Exploration of concatenating string

Q: Which method of concatenating strings has the best performance?

The method on the left has the highest performance.

Plus > String Builder > Join > fmt.Sprintf

However, the performance of String Builder is not consistent (I do not know why yet).
I have observed that it sometimes has worse performance than fmt.Sprintf.

```go
func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "a" + "b"
	}
}

func BenchmarkStringBuilder(b *testing.B) {
  var sb strings.Builder
	for i := 0; i < b.N; i++ {
    sb.WriteString("a")
    sb.WriteString("b")
    _ = sb.String()
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
    _ = strings.Join([]string{"a", "b"}, "")
	}
}

func BenchmarkStringsJoin2(b *testing.B) {
  s := []string{"a", "b"}
	for i := 0; i < b.N; i++ {
    _ = strings.Join(s, "")
	}
}

func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", "a", "b")
	}
}
```

```go
goos: linux
goarch: amd64
pkg: testGO
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkPlus-4                         1000000000               0.2571 ns/op
BenchmarkStringBuilder-4                244353544                8.140 ns/op
BenchmarkStringsJoin-4                  36143734                33.50 ns/op
BenchmarkStringsJoin2-4                 35016460                33.21 ns/op
BenchmarkFmtSprint-4                    13818625                84.10 ns/op
```
