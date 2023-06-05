// Take an existing CPU-bound sequential program, such as the Mandelbrot program
// of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its
// main loop in parallel using channels for communication. How much faster does it
// run on a multiprocessor machine? What is the optimal number of goroutines to use?
package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
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
