package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	a := sort.StringSlice{"a", "b", "c", "b", "a"}
	b := sort.StringSlice{"a", "b", "c", "d", "e"}
	require.True(t, IsPalindrome(a))
	require.False(t, IsPalindrome(b))
}
