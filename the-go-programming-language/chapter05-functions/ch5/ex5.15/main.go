// Write variadic functions max and min, analogous to sum.
// What should these functions do when called with no arguments?
// Write variants that require at least one argument
package main

func max(first int, vals ...int) int {
	max := first
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(first int, vals ...int) int {
	min := first
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
