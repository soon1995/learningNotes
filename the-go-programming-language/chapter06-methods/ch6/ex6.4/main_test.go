package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElem(t *testing.T) {
	m := new(IntSet)
	m.AddAll(1, 2, 3)
	s := m.Elems()
	require.Equal(t, []int{1, 2, 3}, s)
}
