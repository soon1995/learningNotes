// Modify charcount to count letters, digits, and so on in their
// Unicode categories, using functions like unicode.IsLetter
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	Charcount()
}

func Charcount() {
	counts := make(map[rune]int)
	letter := 0
	digit := 0
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsDigit(r) {
			digit++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nletter\tdigit\n")
	fmt.Printf("%d\t%d\n", letter, digit)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characteres\n", invalid)
	}
}
