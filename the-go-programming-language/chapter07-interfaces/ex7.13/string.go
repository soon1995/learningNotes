package main

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	b := &bytes.Buffer{}
	b.WriteString(c.fn)
	b.WriteByte('(')
	for i, c := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(c.String())

	}
	b.WriteByte(')')
	return b.String()
}
