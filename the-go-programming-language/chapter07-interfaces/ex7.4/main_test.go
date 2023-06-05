package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestReader(t *testing.T) {
	buf := &bytes.Buffer{}
	n, err := buf.ReadFrom(NewReader("abcd ef"))
	require.Equal(t, n, int64(7))
	require.NoError(t, err)
}

func TestReaderHTML(t *testing.T) {
	in := `<html>
  <head>
    <title></title>
  </head>
  <body>
    <div>
      <img/>
      <a/>
    </div>
  </body>
  </html>`
	r := NewReader(in)
	n, err := html.Parse(r)
	require.NoError(t, err)
	n = getNextElementNode(n)
	require.Equal(t, "html", n.Data)

	head := getNextElementChild(n)
	require.Equal(t, "head", head.Data)
	require.Equal(t, "title", getNextElementNode(head).Data)

	body := getNextElementSibling(head)
	require.Equal(t, "body", body.Data)

	div := getNextElementNode(body)
	require.Equal(t, "div", div.Data)
	require.Equal(t, "img", getNextElementNode(div).Data)
	require.Equal(t, "a", getNextElementSibling(getNextElementNode(div)).Data)
}

func getNextElementNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		for c.Type == html.ElementNode {
			return c
		}
		res := getNextElementNode(c)
		if res != nil {
			return res
		}
	}
	return nil
}

func getNextElementChild(n *html.Node) *html.Node {
	return getNextElementNode(n)
}

func getNextElementSibling(n *html.Node) *html.Node {
	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			return c
		}
		res := getNextElementSibling(c)
		if res != nil {
			return res
		}
	}
	return nil
}
