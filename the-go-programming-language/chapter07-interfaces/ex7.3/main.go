// Write a String method for the *tree type in gopl.io/ch4/treesort that
// reveals the sequence of values in the tree.
package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	b := &bytes.Buffer{}
	t.string(b)
	return "[" + b.String() + "]"
}

func (t *tree) string(b *bytes.Buffer) {
	if t == nil {
		return
	}
	if t.left != nil {
		t.left.string(b)
		fmt.Fprintf(b, " ")
	}
	fmt.Fprintf(b, "%d", t.value)
	if t.right != nil {
		fmt.Fprintf(b, " ")
		t.right.string(b)
	}
}
