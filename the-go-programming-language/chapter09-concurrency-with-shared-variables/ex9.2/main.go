// Rewrite the PopCount example from Section 2.6.2 so that it initializes the lookup
// table using Sync.Once the first time it is needed. (Realistically, the cost of synchronization would be
// prohibitive for a small and highly optimized function like PopCount.)
package main

import "sync"

// pc[i] is the population count of i.
var pc [256]byte

func initPopCount() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var initPopCountOnce sync.Once

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	initPopCountOnce.Do(initPopCount)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
