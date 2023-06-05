package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := "hello world"
	fmt.Fprintf(c, data)
	require.Equal(t, int64(11), *n)
}
