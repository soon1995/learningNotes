package main

import "testing"

func BenchmarkIntSetAdd(b *testing.B) {
	intset := &IntSet{}
	for i := 0; i < b.N; i++ {
		intset.Add(i)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	m := make(M)
	for i := 0; i < b.N; i++ {
		m.Add(i)
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	intset := &IntSet{}
	intset1 := &IntSet{}
	for i := 0; i < b.N; i++ {
		intset.Add(i)
		intset1.UnionWith(intset)
	}
}

func BenchmarkMapUnionWith(b *testing.B) {
	m := make(M)
	m1 := make(M)
	for i := 0; i < b.N; i++ {
		m.Add(i)
		m1.UnionWIth(m)
	}
}
