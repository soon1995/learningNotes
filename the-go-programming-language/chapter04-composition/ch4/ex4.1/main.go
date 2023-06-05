// Write a function that counts the number of bits that are different in
// two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte{'x'})
	b := sha256.Sum256([]byte{'X'})
	fmt.Println(bitDiff(a[:], b[:]))
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}

	}
	return count
}
