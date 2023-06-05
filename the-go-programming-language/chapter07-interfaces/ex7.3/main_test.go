package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	tr := add(nil, 1)
	add(tr, 2)
	add(tr, 5)
	add(tr, 4)
	add(tr, 3)
  require.Equal(t, "[1 2 3 4 5]", tr.String())
}
