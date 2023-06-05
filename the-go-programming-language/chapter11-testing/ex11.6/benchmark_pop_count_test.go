package main

import (
	"fmt"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
  fmt.Println(PopCountTable(4))
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
  fmt.Println(PopCountShift(4))
	for i := 0; i < b.N; i++ {
		PopCountShift(uint64(i))
	}
}

func BenchmarkPopCountClearRightMost(b *testing.B) {
  fmt.Println(PopCountClearRightMost(4))
	for i := 0; i < b.N; i++ {
		PopCountClearRightMost(uint64(i))
	}
}
