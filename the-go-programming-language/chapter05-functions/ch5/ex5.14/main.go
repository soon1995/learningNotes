// Use the breadthFirst function to explore a different structure. For example,
// you could use the course dependencies from the topoSort example (a directed graph),
// the file system hierarchy on your computer (a tree), or a list of bus or subway routes
// downloaded from your city government's web site (an undirected graph).
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ex5.14 DIR")
	}
	// file system hierarchy
	f := func(item string) []string {
		var list []string
		err := filepath.Walk(item, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			if info.IsDir() {
				list = append(list, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("cannot walk file path: %v", err)
			return nil
		}
		return list
	}
	breadthFirst(f, []string{os.Args[1]})

	// topoSortExample
	// breadthFirst(deps, []string{"programming languages"})
}

// breathFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func deps(course string) []string {
	fmt.Println(course)
	return prereqs[course]
}

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
