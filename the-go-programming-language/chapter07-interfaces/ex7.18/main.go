// Using the token-based decoder API, write a program that will read an arbitrary
// XML document and construct a tree of generic nodes that represents it. Nodes are
// of two kinds: CharData nodes represent text strings, and Element nodes represent named
// elements and their attributes. Each element node has a slice of child nodes.
// You may find the following declaration helpful.
// type Node interface{} // CharData or *Element
// type CharData string
//
//	type Element struct {
//	  Type xml.Name
//	  Attr []xml.Attr
//	  Children []Node
//	}
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	depth := 0
	var fn func(Node)
	fn = func(n1 Node) {
		switch n1 := n1.(type) {
		case *Element:
			fmt.Fprintf(b, "%*s<%s>\n", depth*2, "", n1.Type.Local)
			depth++
			for _, c := range n1.Children {
				fn(c)
			}
			depth--
		case string:
			fmt.Fprintf(b, "%*s%s\n", depth*2, "", n1)
		}
	}
	fn(n)
	return b.String()
}

func Parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []Node
	dummy := &Element{}
	stack = append(stack, dummy)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch x := tok.(type) {
		case xml.StartElement:
			el := &Element{
				Type: x.Name,
				Attr: x.Attr,
			}
			err := addChild(stack[len(stack)-1], el)
			if err != nil {
				return nil, err
			}
			stack = append(stack, el)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(bytes.TrimSpace(x)) == 0 {
				continue
			}
			err := addChild(stack[len(stack)-1], string(x))
			if err != nil {
				return nil, err
			}
		}
	}
	return dummy.Children[0], nil
}

func addChild(p, c Node) error {
	if v, ok := p.(*Element); ok {
		v.Children = append(v.Children, c)
		return nil
	}
	return fmt.Errorf("unexpected type %T: %v", p, p)
}

var file = flag.String("f", "", "file")

func main() {
	flag.Parse()
	var in io.Reader
	in = os.Stdin
	if *file != "" {
		buf := &bytes.Buffer{}
		fl, err := os.Open(*file)
		if err != nil {
			log.Fatalf("cannot open file: %s", err)
		}
		_, err = io.Copy(buf, fl)
		if err != nil {
			log.Fatalf("cannot copy buffer: %s", err)
		}
		if err := fl.Close(); err != nil {
			log.Fatalf("cannot close file: %s", err)
		}
		in = buf
	}
	n, err := Parse(in)
	if err != nil {
		log.Fatalf("cannot parse: %s", err)
	}
	fmt.Println(n)
}

// cat test.xml | go run .
