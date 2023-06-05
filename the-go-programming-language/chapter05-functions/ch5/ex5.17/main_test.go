package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	in := `
  <html>
    <head>
    </head>
    <body>
      <img/>
      <h1></h1>
      <h2></h2>
      <h3></h3>
      <h4></h4>
    </body>
  </html>
  `
  r :=strings.NewReader(in)
  doc, err := html.Parse(r)
  require.NoError(t, err)
  require.Equal(t, 1, len(ElementsByTagName(doc, "img")))
  require.Equal(t, 4, len(ElementsByTagName(doc, "h1", "h2", "h3", "h4")))
}
