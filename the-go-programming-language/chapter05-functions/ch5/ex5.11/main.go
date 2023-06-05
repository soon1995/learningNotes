// The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string, key string)
	visitAll = func(n map[string][]string, key string) {
		for _, item := range n[key] {
			if !seen[item] {
				if contain(n[item], key) {
					fmt.Printf("cycle dependency: %s -> %s\n", key, item)
				}
				seen[item] = true
				visitAll(m, item)
				order = append(order, item)
			}
		}
		if !seen[key] {
			seen[key] = true
			visitAll(m, key)
			order = append(order, key)
		}
	}
	for k := range m {
		visitAll(m, k)
	}
	return order
}

func contain(s []string, target string) (exist bool) {
	for _, v := range s {
		if target == v {
			exist = true
		}
	}
	return
}
