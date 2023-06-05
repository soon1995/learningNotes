// Construct a pipeline that connects an arbitrary number of goroutines with
// channels. What is the maximum number of pipeline stages you can create without
// running out of memory? How long does a value take to transit the entire pipeline?
// copied torbiak/gopl/ex9.4/pipe.go
// my : 3400000 76s
package main

import "testing"

func BenchmarkPipeline(b *testing.B) {
	in, out := pipeline(3400000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
