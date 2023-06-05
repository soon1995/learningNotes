package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterceptWith(t *testing.T){
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(9, 42)
  x.IntersectWith(&y)
  fmt.Println(&x)
  require.Equal(t, 1, x.Len())
  require.True(t, x.Has(9))
}

func TestDifferentWith(t *testing.T){
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(9, 42,100,1000)
  x.DifferenceWith(&y)
  require.Equal(t, 2, x.Len())
  require.True(t, x.Has(1))
  require.True(t, x.Has(144))
}

func TestSymmetryDifference(t *testing.T){
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(9, 42,100,1000)
  x.SymmetricDifference(&y)
  require.Equal(t, 5, x.Len())
  require.True(t, x.Has(1))
  require.True(t, x.Has(42))
  require.True(t, x.Has(100))
  require.True(t, x.Has(144))
  require.True(t, x.Has(1000))
}
