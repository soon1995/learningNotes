// Measure how the performance of a compute-bound parallel program (see Exercise 8.5)
// varies with GOMAXPROCS. What is the optimal value on your computer? How many
// CPU does your computer have?
package main

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

func main() {
  fmt.Println(runtime.NumCPU())
  fmt.Println(runtime.GOMAXPROCS(0))
	t1 := time.Now()
	SurfaceNoOptimized(io.Discard)
	fmt.Printf("Surface not optimized: %s\n", time.Now().Sub(t1))
	t2 := time.Now()
	SurfaceOptimized(io.Discard)
	fmt.Printf("Surface optimized: %s\n", time.Now().Sub(t2))
	t3 := time.Now()
	SurfaceOptimized2(io.Discard)
	fmt.Printf("Surface optimized: %s\n", time.Now().Sub(t3))
}
