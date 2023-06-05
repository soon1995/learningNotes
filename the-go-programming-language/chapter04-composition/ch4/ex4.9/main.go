// Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before
// the first call to Scan to break the input into words instead of lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

func wordfreq() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordCount[word]++
	}
	if input.Err() != nil {
		fmt.Fprintln(os.Stderr, input.Err())
		os.Exit(1)
	}
	for k, v := range wordCount {
		fmt.Printf("%-30s %d\n", k, v)
	}
}
