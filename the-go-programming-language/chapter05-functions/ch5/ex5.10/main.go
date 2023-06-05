// Rewrite topoSort to use maps instead of slices and eliminate the initial sort.
// Verify that the results, though nondeterministric, are valid topological orderings.
package main

import "fmt"

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string, key string)
	visitAll = func(n map[string][]string, key string) {
		for _, item := range n[key] {
			if !seen[item] {
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

func main() {
	a := topoSort(prereqs)
	for _, v := range a {
		fmt.Println(v)
	}
}
