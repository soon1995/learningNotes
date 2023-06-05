// Construct a pipeline that connects an arbitrary number of goroutines with
// channels. What is the maximum number of pipeline stages you can create without
// running out of memory? How long does a value take to transit the entire pipeline?
// copied torbiak/gopl/ex9.4/pipe.go
package main

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}

func main() {
	in, out := pipeline(3400000)
  in <- 1
  <-out
  close(in)
}
