// Extend xmlselect so that elements may be selected not just by name, but
// by their attributes too, in the manner of CSS, so that, for instance, an
// element like <div id="page" class="wide"> could be selected by a matching
// id or class as well as its named
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var targets map[string]map[string]bool

// cat test.xml | go run . a class=testa id=b
func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	targets = make(map[string]map[string]bool)
	if len(os.Args) > 2 {
		for _, arg := range os.Args[2:] {
			fields := strings.Split(arg, "=")
			if len(fields) != 2 {
				log.Fatalf("bad arguments: %s, expect key=value", arg)
			}
			if targets[fields[0]] == nil {
				targets[fields[0]] = make(map[string]bool)
			}
			targets[fields[0]][fields[1]] = true
		}
	}
	dec := xml.NewDecoder(os.Stdin)
	var stack []interface{}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("xmlselect %v", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			if len(tok.Attr) > 0 {
				stack = append(stack, tok.Attr)
			}
		case xml.EndElement:
			for len(stack) > 0 {
				el := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if _, ok := el.(string); ok {
					break
				}
			}
		case xml.CharData:
			if isElement(stack, os.Args[1]) && containsAllAttribute(stack, targets) {
				print(stack, tok)
			}
		}
	}
}

func isElement(stack []interface{}, el string) bool {
	for i := len(stack) - 1; i >= 0; i-- {
		switch x := stack[i].(type) {
		case string:
			return x == el
		}
	}
	return false
}

func containsAllAttribute(stack []interface{}, m map[string]map[string]bool) bool {
	unmatched := totalKeys(m)
	if unmatched == 0 {
		return true
	}
	for i := len(stack) - 1; i >= 0; i-- {
		switch x := stack[i].(type) {
		case string:
			return false
		case []xml.Attr:
			for _, v := range x {
				if _, ok := m[v.Name.Local][v.Value]; ok {
					unmatched--
				}
				if unmatched == 0 {
					return true
				}
			}
		}
	}
	return false
}

func totalKeys(m map[string]map[string]bool) (c int) {
	for _, v := range m {
		c += len(v)
	}
	return
}

func print(stack []interface{}, chardata xml.CharData) {
	b := &bytes.Buffer{}
	for i, v := range stack {
		if i != 0 {
			b.WriteByte(' ')
		}
		switch x := v.(type) {
		case string:
			b.WriteString(x)
		case []xml.Attr:
			b.WriteString("[")
			for i, v := range x {
				if i == 0 {
					b.WriteByte(' ')
				}
				b.WriteString(fmt.Sprintf("%s=%s ", v.Name.Local, v.Value))
			}
			b.WriteString("]")
		}
	}
	b.WriteString(": ")
	b.Write(chardata)
	fmt.Println(b.String())
}
