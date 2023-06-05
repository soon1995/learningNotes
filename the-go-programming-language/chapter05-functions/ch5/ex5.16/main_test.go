package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	require.Equal(t, "asasa", join("s", "a", "a", "a"))
	require.Equal(t, "a\ta\ta", join("\t", "a", "a", "a"))
	require.Equal(t, "", join("\t"))
}
