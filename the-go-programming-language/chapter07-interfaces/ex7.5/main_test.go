package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLimitReader(t *testing.T) {
	in := strings.NewReader("hello world should become hello world")
	r := &MyLimitReader{in, 11}
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(r)
	require.Error(t, io.EOF, err)
	require.Equal(t, int64(11), n)

	in = strings.NewReader("hello world")
	r = &MyLimitReader{in, 11}
	n, err = b.ReadFrom(r)
	require.NoError(t, err)
	require.Equal(t, int64(11), n)
}
