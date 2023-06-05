package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpand(t *testing.T) {
  subs := make(map[string]string, 0)
  subs["foo"] = "bar"
  subs["hello"] = "world"

  f := func(s string) string {
    v, ok := subs[s]
    if !ok {
      return "$" + s
    }
    return v
  }
  tc1 := "$foo is handsome"
  expect1 := "bar is handsome"
  got1 := expand(tc1, f)
  require.Equal(t, expect1, got1)

  tc2 := "$hello is $handsome"
  expect2 := "world is $handsome"
  got2 := expand(tc2, f)
  require.Equal(t, expect2, got2)
}
