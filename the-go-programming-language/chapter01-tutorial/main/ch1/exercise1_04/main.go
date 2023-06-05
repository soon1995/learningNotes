package main

import (
	"bufio"
	"fmt"
	"os"
)

// Modify dup2 to print the names of all files in which each duplicated
// line occurs
func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fileMap := range counts {
		total := 0
		filenames := make([]string, 0)
		for filename, count := range fileMap {
			filenames = append(filenames, filename)
			total += count
		}
		if total > 1 {
			fmt.Printf("%d\t%s\t%v\n", total, line, filenames)
		}

	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
