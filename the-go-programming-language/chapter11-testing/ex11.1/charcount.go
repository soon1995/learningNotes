// Write tests for the charcount program in Section 4.3
package charcount

import (
	"unicode"
	"unicode/utf8"
)

func Charcount(in string) (counts map[rune]int, utflen [4]int) {
	counts = make(map[rune]int)     // counts of Unicode characters
	for _, r := range in {
		n := utf8.RuneLen(r)
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return
}
