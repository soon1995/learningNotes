package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	require.Equal(t, 5, max(1, 2, 3, 4, 5))
}
func TestMin(t *testing.T) {
	require.Equal(t, 1, min(1, 2, 3, 4, 5))
}
