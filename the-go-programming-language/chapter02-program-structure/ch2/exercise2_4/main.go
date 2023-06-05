// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
// bit positions, testing the rightmost bit each time. Compare its performance to the table 
// lookup version.
package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func main() {
  t1 := time.Now()
  PopCountSingleExpression(99999999999)
	fmt.Println(time.Now().Sub(t1)) // 50ns
  t2 := time.Now()
  PopCountLoop(99999999999)
	fmt.Println(time.Now().Sub(t2)) // 80ns
  t3 := time.Now()
  PopCountShift(99999999999)
	fmt.Println(time.Now().Sub(t3)) // 90ns
}

func init() {
	for i := range pc {
		// pc[i] is the population count of i
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	res := byte(0)
	for i := 0; i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}

func PopCountSingleExpression(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int {
  n := 0
  for i := uint(0); i < 64; i ++ {
    if x&1 == 1 {
      n++
    }
    x = x >> 1
  }
  return n
}
